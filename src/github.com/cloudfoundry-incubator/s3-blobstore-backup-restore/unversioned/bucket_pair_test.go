package unversioned_test

import (
	"github.com/cloudfoundry-incubator/s3-blobstore-backup-restore/s3/fakes"
	"github.com/cloudfoundry-incubator/s3-blobstore-backup-restore/unversioned"

	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BucketPair", func() {
	var (
		liveBucket   *fakes.FakeUnversionedBucket
		backupBucket *fakes.FakeUnversionedBucket
		bucketPair   unversioned.S3BucketPair
		address      unversioned.BackupBucketAddress
		err          error
	)

	BeforeEach(func() {
		liveBucket = new(fakes.FakeUnversionedBucket)
		backupBucket = new(fakes.FakeUnversionedBucket)
		bucketPair = unversioned.NewS3BucketPair(liveBucket, backupBucket)

		liveBucket.NameReturns("liveBucket")
		liveBucket.RegionReturns("liveBucketRegion")
		backupBucket.NameReturns("backupBucket")
		backupBucket.RegionReturns("backupBucketRegion")
	})

	Describe("Backup", func() {
		JustBeforeEach(func() {
			address, err = bucketPair.Backup("destination-string")
		})

		Context("When there are files in the bucket", func() {
			BeforeEach(func() {
				liveBucket.ListFilesReturns([]string{"path1/file1", "path2/file2"}, nil)
			})

			It("copies all the files in the bucket", func() {
				By("succeeding", func() {
					Expect(err).NotTo(HaveOccurred())
				})
				By("Listing the files in the bucket", func() {
					Expect(liveBucket.ListFilesCallCount()).To(Equal(1))
				})

				By("calling copy for each file in the bucket", func() {
					Expect(backupBucket.CopyObjectCallCount()).To(Equal(2))

					var keys []string

					key, originPath, destinationPath, originBucketName, originBucketRegion := backupBucket.CopyObjectArgsForCall(1)
					Expect(originPath).To(Equal(""))
					Expect(destinationPath).To(Equal("destination-string"))
					Expect(originBucketName).To(Equal("liveBucket"))
					Expect(originBucketRegion).To(Equal("liveBucketRegion"))
					keys = append(keys, key)

					key, originPath, destinationPath, originBucketName, originBucketRegion = backupBucket.CopyObjectArgsForCall(0)
					Expect(originPath).To(Equal(""))
					Expect(destinationPath).To(Equal("destination-string"))
					Expect(originBucketName).To(Equal("liveBucket"))
					Expect(originBucketRegion).To(Equal("liveBucketRegion"))
					keys = append(keys, key)

					Expect(keys).To(ConsistOf("path1/file1", "path2/file2"))
				})

				By("returning the address of the backup bucket", func() {
					Expect(address).To(Equal(unversioned.BackupBucketAddress{
						BucketName:   "backupBucket",
						BucketRegion: "backupBucketRegion",
						Path:         "destination-string",
						EmptyBackup:  false,
					}))
				})
			})

			Context("when CopyObject fails", func() {
				BeforeEach(func() {
					backupBucket.CopyObjectReturnsOnCall(0, fmt.Errorf("cannot copy first file"))
					backupBucket.CopyObjectReturnsOnCall(1, fmt.Errorf("cannot copy second file"))
				})

				It("should fail", func() {
					Expect(err).To(MatchError(ContainSubstring("cannot copy first file")))
					Expect(err).To(MatchError(ContainSubstring("cannot copy second file")))
				})
			})
		})

		Context("When there are no files in the bucket", func() {
			BeforeEach(func() {
				liveBucket.ListFilesReturns([]string{}, nil)
			})

			It("Records that information in the backup artifact", func() {

				By("succeeding", func() {
					Expect(err).NotTo(HaveOccurred())
				})

				By("not calling copy", func() {
					Expect(backupBucket.CopyObjectCallCount()).To(Equal(0))
				})

				By("recording that the backup was empty", func() {
					Expect(address).To(Equal(unversioned.BackupBucketAddress{
						BucketName:   "backupBucket",
						BucketRegion: "backupBucketRegion",
						Path:         "destination-string",
						EmptyBackup:  true,
					}))
				})
			})

		})

		Context("when ListFiles fails", func() {
			BeforeEach(func() {
				liveBucket.ListFilesReturns([]string{}, fmt.Errorf("cannot list files"))
			})

			It("should fail", func() {
				Expect(err).To(MatchError("cannot list files"))
			})
		})

	})

	Describe("Restore", func() {
		JustBeforeEach(func() {
			err = bucketPair.Restore("2015-12-13-05-06-07/my_bucket")
		})

		BeforeEach(func() {
			backupBucket.ListFilesReturns([]string{"my_key", "another_key"}, nil)
		})

		It("successfully copies from the backup bucket to the live bucket", func() {
			By("not returning an error", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			By("copying from the backup location to the live location", func() {
				Expect(backupBucket.ListFilesCallCount()).To(Equal(1))
				Expect(backupBucket.ListFilesArgsForCall(0)).To(Equal("2015-12-13-05-06-07/my_bucket"))

				Expect(liveBucket.CopyObjectCallCount()).To(Equal(2))

				var keys []string

				key, originPath, destinationPath, originBucketName, originBucketRegion := liveBucket.CopyObjectArgsForCall(0)
				Expect(originPath).To(Equal("2015-12-13-05-06-07/my_bucket"))
				Expect(destinationPath).To(Equal(""))
				Expect(originBucketName).To(Equal(backupBucket.Name()))
				Expect(originBucketRegion).To(Equal(backupBucket.Region()))
				keys = append(keys, key)

				key, originPath, destinationPath, originBucketName, originBucketRegion = liveBucket.CopyObjectArgsForCall(1)
				Expect(originPath).To(Equal("2015-12-13-05-06-07/my_bucket"))
				Expect(destinationPath).To(Equal(""))
				Expect(originBucketName).To(Equal(backupBucket.Name()))
				Expect(originBucketRegion).To(Equal(backupBucket.Region()))
				keys = append(keys, key)

				Expect(keys).To(ConsistOf("my_key", "another_key"))
			})
		})

		Context("When ListFiles errors", func() {
			BeforeEach(func() {
				backupBucket.ListFilesReturns([]string{}, fmt.Errorf("cannot list files"))
			})

			It("errors", func() {
				Expect(err).To(MatchError(ContainSubstring("cannot list files")))
			})
		})

		Context("When CopyObject errors", func() {
			BeforeEach(func() {
				liveBucket.CopyObjectReturns(fmt.Errorf("cannot copy object"))
			})

			It("errors", func() {
				Expect(err).To(MatchError(ContainSubstring("cannot copy object")))
			})
		})

		Context("When listFiles returns no files", func() {
			BeforeEach(func() {
				backupBucket.ListFilesReturns([]string{}, nil)
			})

			It("errors saying there should be a backup", func() {
				Expect(err).To(MatchError(ContainSubstring("no files found in 2015-12-13-05-06-07/my_bucket " +
					"in bucket backupBucket to restore")))
			})
		})
	})

	Describe("CheckValidity", func() {
		Context("when the live bucket and the backup bucket are not the same", func() {
			It("returns nil", func() {
				Expect(unversioned.NewS3BucketPair(liveBucket, backupBucket).CheckValidity()).To(BeNil())
			})
		})

		Context("when the live bucket and the backup bucket are the same", func() {
			It("returns an error", func() {
				Expect(unversioned.NewS3BucketPair(liveBucket, liveBucket).CheckValidity()).To(MatchError("live bucket and backup bucket cannot be the same"))
			})
		})
	})

	Describe("LiveBucketName", func() {
		It("returns the name of the live bucket", func() {
			Expect(bucketPair.LiveBucketName()).To(Equal(liveBucket.Name()))
		})
	})

	Describe("BackupBucketName", func() {
		It("returns the name of the backup bucket", func() {
			Expect(bucketPair.BackupBucketName()).To(Equal(backupBucket.Name()))
		})
	})
})
