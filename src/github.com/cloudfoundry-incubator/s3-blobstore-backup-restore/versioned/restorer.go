package versioned

import (
	"fmt"

	"github.com/cloudfoundry-incubator/s3-blobstore-backup-restore/s3bucket"
)

type Restorer struct {
	destinationBuckets map[string]s3bucket.VersionedBucket
	sourceArtifact     Artifact
}

func NewRestorer(destinationBuckets map[string]s3bucket.VersionedBucket, sourceArtifact Artifact) Restorer {
	return Restorer{destinationBuckets: destinationBuckets, sourceArtifact: sourceArtifact}
}

func (r Restorer) Run() error {
	bucketSnapshots, err := r.sourceArtifact.Load()
	if err != nil {
		return err
	}

	for identifier := range bucketSnapshots {
		_, exists := r.destinationBuckets[identifier]

		if !exists {
			return fmt.Errorf("no entry found in restore config for bucket: %s", identifier)
		}
	}

	for identifier, destinationBucket := range r.destinationBuckets {
		bucketSnapshot, exists := bucketSnapshots[identifier]

		if !exists {
			return fmt.Errorf("no entry found in backup artifact for bucket: %s", identifier)
		}

		err := destinationBucket.CheckIfVersioned()
		if err != nil {
			return err
		}

		for _, versionToCopy := range bucketSnapshot.Versions {
			err = destinationBucket.CopyVersion(
				versionToCopy.BlobKey,
				versionToCopy.Id,
				bucketSnapshot.BucketName,
				bucketSnapshot.RegionName,
			)

			if err != nil {
				return err
			}
		}
	}

	return nil
}
