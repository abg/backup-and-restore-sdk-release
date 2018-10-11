// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	gcs "github.com/cloudfoundry-incubator/gcs-blobstore-backup-restore"
)

type FakeBucket struct {
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct{}
	nameReturns     struct {
		result1 string
	}
	nameReturnsOnCall map[int]struct {
		result1 string
	}
	ListBlobsStub        func() ([]gcs.Blob, error)
	listBlobsMutex       sync.RWMutex
	listBlobsArgsForCall []struct{}
	listBlobsReturns     struct {
		result1 []gcs.Blob
		result2 error
	}
	listBlobsReturnsOnCall map[int]struct {
		result1 []gcs.Blob
		result2 error
	}
	CopyBlobStub        func(string, string) (int64, error)
	copyBlobMutex       sync.RWMutex
	copyBlobArgsForCall []struct {
		arg1 string
		arg2 string
	}
	copyBlobReturns struct {
		result1 int64
		result2 error
	}
	copyBlobReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBucket) Name() string {
	fake.nameMutex.Lock()
	ret, specificReturn := fake.nameReturnsOnCall[len(fake.nameArgsForCall)]
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct{}{})
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.nameReturns.result1
}

func (fake *FakeBucket) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeBucket) NameReturns(result1 string) {
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeBucket) NameReturnsOnCall(i int, result1 string) {
	fake.NameStub = nil
	if fake.nameReturnsOnCall == nil {
		fake.nameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.nameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeBucket) ListBlobs() ([]gcs.Blob, error) {
	fake.listBlobsMutex.Lock()
	ret, specificReturn := fake.listBlobsReturnsOnCall[len(fake.listBlobsArgsForCall)]
	fake.listBlobsArgsForCall = append(fake.listBlobsArgsForCall, struct{}{})
	fake.recordInvocation("ListBlobs", []interface{}{})
	fake.listBlobsMutex.Unlock()
	if fake.ListBlobsStub != nil {
		return fake.ListBlobsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listBlobsReturns.result1, fake.listBlobsReturns.result2
}

func (fake *FakeBucket) ListBlobsCallCount() int {
	fake.listBlobsMutex.RLock()
	defer fake.listBlobsMutex.RUnlock()
	return len(fake.listBlobsArgsForCall)
}

func (fake *FakeBucket) ListBlobsReturns(result1 []gcs.Blob, result2 error) {
	fake.ListBlobsStub = nil
	fake.listBlobsReturns = struct {
		result1 []gcs.Blob
		result2 error
	}{result1, result2}
}

func (fake *FakeBucket) ListBlobsReturnsOnCall(i int, result1 []gcs.Blob, result2 error) {
	fake.ListBlobsStub = nil
	if fake.listBlobsReturnsOnCall == nil {
		fake.listBlobsReturnsOnCall = make(map[int]struct {
			result1 []gcs.Blob
			result2 error
		})
	}
	fake.listBlobsReturnsOnCall[i] = struct {
		result1 []gcs.Blob
		result2 error
	}{result1, result2}
}

func (fake *FakeBucket) CopyBlob(arg1 string, arg2 string) (int64, error) {
	fake.copyBlobMutex.Lock()
	ret, specificReturn := fake.copyBlobReturnsOnCall[len(fake.copyBlobArgsForCall)]
	fake.copyBlobArgsForCall = append(fake.copyBlobArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("CopyBlob", []interface{}{arg1, arg2})
	fake.copyBlobMutex.Unlock()
	if fake.CopyBlobStub != nil {
		return fake.CopyBlobStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.copyBlobReturns.result1, fake.copyBlobReturns.result2
}

func (fake *FakeBucket) CopyBlobCallCount() int {
	fake.copyBlobMutex.RLock()
	defer fake.copyBlobMutex.RUnlock()
	return len(fake.copyBlobArgsForCall)
}

func (fake *FakeBucket) CopyBlobArgsForCall(i int) (string, string) {
	fake.copyBlobMutex.RLock()
	defer fake.copyBlobMutex.RUnlock()
	return fake.copyBlobArgsForCall[i].arg1, fake.copyBlobArgsForCall[i].arg2
}

func (fake *FakeBucket) CopyBlobReturns(result1 int64, result2 error) {
	fake.CopyBlobStub = nil
	fake.copyBlobReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeBucket) CopyBlobReturnsOnCall(i int, result1 int64, result2 error) {
	fake.CopyBlobStub = nil
	if fake.copyBlobReturnsOnCall == nil {
		fake.copyBlobReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.copyBlobReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeBucket) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.listBlobsMutex.RLock()
	defer fake.listBlobsMutex.RUnlock()
	fake.copyBlobMutex.RLock()
	defer fake.copyBlobMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBucket) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ gcs.Bucket = new(FakeBucket)
