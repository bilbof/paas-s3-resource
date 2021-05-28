// This file was generated by counterfeiter
package fakes

import (
	"sync"

	s3resource "github.com/alphagov/paas-s3-resource"
)

type FakeS3Client struct {
	BucketFilesStub        func(bucketName string, prefixHint string) ([]string, error)
	bucketFilesMutex       sync.RWMutex
	bucketFilesArgsForCall []struct {
		bucketName string
		prefixHint string
	}
	bucketFilesReturns struct {
		result1 []string
		result2 error
	}
	BucketFileVersionsStub        func(bucketName string, remotePath string) ([]string, error)
	bucketFileVersionsMutex       sync.RWMutex
	bucketFileVersionsArgsForCall []struct {
		bucketName string
		remotePath string
	}
	bucketFileVersionsReturns struct {
		result1 []string
		result2 error
	}
	UploadFileStub        func(bucketName string, remotePath string, localPath string, options s3resource.UploadFileOptions) (string, error)
	uploadFileMutex       sync.RWMutex
	uploadFileArgsForCall []struct {
		bucketName string
		remotePath string
		localPath  string
		options    s3resource.UploadFileOptions
	}
	uploadFileReturns struct {
		result1 string
		result2 error
	}
	DownloadFileStub        func(bucketName string, remotePath string, versionID string, localPath string) error
	downloadFileMutex       sync.RWMutex
	downloadFileArgsForCall []struct {
		bucketName string
		remotePath string
		versionID  string
		localPath  string
	}
	downloadFileReturns struct {
		result1 error
	}
	DeleteFileStub        func(bucketName string, remotePath string) error
	deleteFileMutex       sync.RWMutex
	deleteFileArgsForCall []struct {
		bucketName string
		remotePath string
	}
	deleteFileReturns struct {
		result1 error
	}
	DeleteVersionedFileStub        func(bucketName string, remotePath string, versionID string) error
	deleteVersionedFileMutex       sync.RWMutex
	deleteVersionedFileArgsForCall []struct {
		bucketName string
		remotePath string
		versionID  string
	}
	deleteVersionedFileReturns struct {
		result1 error
	}
	URLStub        func(bucketName string, remotePath string, private bool, versionID string) string
	uRLMutex       sync.RWMutex
	uRLArgsForCall []struct {
		bucketName string
		remotePath string
		private    bool
		versionID  string
	}
	uRLReturns struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeS3Client) BucketFiles(bucketName string, prefixHint string) ([]string, error) {
	fake.bucketFilesMutex.Lock()
	fake.bucketFilesArgsForCall = append(fake.bucketFilesArgsForCall, struct {
		bucketName string
		prefixHint string
	}{bucketName, prefixHint})
	fake.recordInvocation("BucketFiles", []interface{}{bucketName, prefixHint})
	fake.bucketFilesMutex.Unlock()
	if fake.BucketFilesStub != nil {
		return fake.BucketFilesStub(bucketName, prefixHint)
	} else {
		return fake.bucketFilesReturns.result1, fake.bucketFilesReturns.result2
	}
}

func (fake *FakeS3Client) BucketFilesCallCount() int {
	fake.bucketFilesMutex.RLock()
	defer fake.bucketFilesMutex.RUnlock()
	return len(fake.bucketFilesArgsForCall)
}

func (fake *FakeS3Client) BucketFilesArgsForCall(i int) (string, string) {
	fake.bucketFilesMutex.RLock()
	defer fake.bucketFilesMutex.RUnlock()
	return fake.bucketFilesArgsForCall[i].bucketName, fake.bucketFilesArgsForCall[i].prefixHint
}

func (fake *FakeS3Client) BucketFilesReturns(result1 []string, result2 error) {
	fake.BucketFilesStub = nil
	fake.bucketFilesReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeS3Client) BucketFileVersions(bucketName string, remotePath string) ([]string, error) {
	fake.bucketFileVersionsMutex.Lock()
	fake.bucketFileVersionsArgsForCall = append(fake.bucketFileVersionsArgsForCall, struct {
		bucketName string
		remotePath string
	}{bucketName, remotePath})
	fake.recordInvocation("BucketFileVersions", []interface{}{bucketName, remotePath})
	fake.bucketFileVersionsMutex.Unlock()
	if fake.BucketFileVersionsStub != nil {
		return fake.BucketFileVersionsStub(bucketName, remotePath)
	} else {
		return fake.bucketFileVersionsReturns.result1, fake.bucketFileVersionsReturns.result2
	}
}

func (fake *FakeS3Client) BucketFileVersionsCallCount() int {
	fake.bucketFileVersionsMutex.RLock()
	defer fake.bucketFileVersionsMutex.RUnlock()
	return len(fake.bucketFileVersionsArgsForCall)
}

func (fake *FakeS3Client) BucketFileVersionsArgsForCall(i int) (string, string) {
	fake.bucketFileVersionsMutex.RLock()
	defer fake.bucketFileVersionsMutex.RUnlock()
	return fake.bucketFileVersionsArgsForCall[i].bucketName, fake.bucketFileVersionsArgsForCall[i].remotePath
}

func (fake *FakeS3Client) BucketFileVersionsReturns(result1 []string, result2 error) {
	fake.BucketFileVersionsStub = nil
	fake.bucketFileVersionsReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeS3Client) UploadFile(bucketName string, remotePath string, localPath string, options s3resource.UploadFileOptions) (string, error) {
	fake.uploadFileMutex.Lock()
	fake.uploadFileArgsForCall = append(fake.uploadFileArgsForCall, struct {
		bucketName string
		remotePath string
		localPath  string
		options    s3resource.UploadFileOptions
	}{bucketName, remotePath, localPath, options})
	fake.recordInvocation("UploadFile", []interface{}{bucketName, remotePath, localPath, options})
	fake.uploadFileMutex.Unlock()
	if fake.UploadFileStub != nil {
		return fake.UploadFileStub(bucketName, remotePath, localPath, options)
	} else {
		return fake.uploadFileReturns.result1, fake.uploadFileReturns.result2
	}
}

func (fake *FakeS3Client) UploadFileCallCount() int {
	fake.uploadFileMutex.RLock()
	defer fake.uploadFileMutex.RUnlock()
	return len(fake.uploadFileArgsForCall)
}

func (fake *FakeS3Client) UploadFileArgsForCall(i int) (string, string, string, s3resource.UploadFileOptions) {
	fake.uploadFileMutex.RLock()
	defer fake.uploadFileMutex.RUnlock()
	return fake.uploadFileArgsForCall[i].bucketName, fake.uploadFileArgsForCall[i].remotePath, fake.uploadFileArgsForCall[i].localPath, fake.uploadFileArgsForCall[i].options
}

func (fake *FakeS3Client) UploadFileReturns(result1 string, result2 error) {
	fake.UploadFileStub = nil
	fake.uploadFileReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeS3Client) DownloadFile(bucketName string, remotePath string, versionID string, localPath string) error {
	fake.downloadFileMutex.Lock()
	fake.downloadFileArgsForCall = append(fake.downloadFileArgsForCall, struct {
		bucketName string
		remotePath string
		versionID  string
		localPath  string
	}{bucketName, remotePath, versionID, localPath})
	fake.recordInvocation("DownloadFile", []interface{}{bucketName, remotePath, versionID, localPath})
	fake.downloadFileMutex.Unlock()
	if fake.DownloadFileStub != nil {
		return fake.DownloadFileStub(bucketName, remotePath, versionID, localPath)
	} else {
		return fake.downloadFileReturns.result1
	}
}

func (fake *FakeS3Client) DownloadFileCallCount() int {
	fake.downloadFileMutex.RLock()
	defer fake.downloadFileMutex.RUnlock()
	return len(fake.downloadFileArgsForCall)
}

func (fake *FakeS3Client) DownloadFileArgsForCall(i int) (string, string, string, string) {
	fake.downloadFileMutex.RLock()
	defer fake.downloadFileMutex.RUnlock()
	return fake.downloadFileArgsForCall[i].bucketName, fake.downloadFileArgsForCall[i].remotePath, fake.downloadFileArgsForCall[i].versionID, fake.downloadFileArgsForCall[i].localPath
}

func (fake *FakeS3Client) DownloadFileReturns(result1 error) {
	fake.DownloadFileStub = nil
	fake.downloadFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeS3Client) DeleteFile(bucketName string, remotePath string) error {
	fake.deleteFileMutex.Lock()
	fake.deleteFileArgsForCall = append(fake.deleteFileArgsForCall, struct {
		bucketName string
		remotePath string
	}{bucketName, remotePath})
	fake.recordInvocation("DeleteFile", []interface{}{bucketName, remotePath})
	fake.deleteFileMutex.Unlock()
	if fake.DeleteFileStub != nil {
		return fake.DeleteFileStub(bucketName, remotePath)
	} else {
		return fake.deleteFileReturns.result1
	}
}

func (fake *FakeS3Client) DeleteFileCallCount() int {
	fake.deleteFileMutex.RLock()
	defer fake.deleteFileMutex.RUnlock()
	return len(fake.deleteFileArgsForCall)
}

func (fake *FakeS3Client) DeleteFileArgsForCall(i int) (string, string) {
	fake.deleteFileMutex.RLock()
	defer fake.deleteFileMutex.RUnlock()
	return fake.deleteFileArgsForCall[i].bucketName, fake.deleteFileArgsForCall[i].remotePath
}

func (fake *FakeS3Client) DeleteFileReturns(result1 error) {
	fake.DeleteFileStub = nil
	fake.deleteFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeS3Client) DeleteVersionedFile(bucketName string, remotePath string, versionID string) error {
	fake.deleteVersionedFileMutex.Lock()
	fake.deleteVersionedFileArgsForCall = append(fake.deleteVersionedFileArgsForCall, struct {
		bucketName string
		remotePath string
		versionID  string
	}{bucketName, remotePath, versionID})
	fake.recordInvocation("DeleteVersionedFile", []interface{}{bucketName, remotePath, versionID})
	fake.deleteVersionedFileMutex.Unlock()
	if fake.DeleteVersionedFileStub != nil {
		return fake.DeleteVersionedFileStub(bucketName, remotePath, versionID)
	} else {
		return fake.deleteVersionedFileReturns.result1
	}
}

func (fake *FakeS3Client) DeleteVersionedFileCallCount() int {
	fake.deleteVersionedFileMutex.RLock()
	defer fake.deleteVersionedFileMutex.RUnlock()
	return len(fake.deleteVersionedFileArgsForCall)
}

func (fake *FakeS3Client) DeleteVersionedFileArgsForCall(i int) (string, string, string) {
	fake.deleteVersionedFileMutex.RLock()
	defer fake.deleteVersionedFileMutex.RUnlock()
	return fake.deleteVersionedFileArgsForCall[i].bucketName, fake.deleteVersionedFileArgsForCall[i].remotePath, fake.deleteVersionedFileArgsForCall[i].versionID
}

func (fake *FakeS3Client) DeleteVersionedFileReturns(result1 error) {
	fake.DeleteVersionedFileStub = nil
	fake.deleteVersionedFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeS3Client) URL(bucketName string, remotePath string, private bool, versionID string) string {
	fake.uRLMutex.Lock()
	fake.uRLArgsForCall = append(fake.uRLArgsForCall, struct {
		bucketName string
		remotePath string
		private    bool
		versionID  string
	}{bucketName, remotePath, private, versionID})
	fake.recordInvocation("URL", []interface{}{bucketName, remotePath, private, versionID})
	fake.uRLMutex.Unlock()
	if fake.URLStub != nil {
		return fake.URLStub(bucketName, remotePath, private, versionID)
	} else {
		return fake.uRLReturns.result1
	}
}

func (fake *FakeS3Client) URLCallCount() int {
	fake.uRLMutex.RLock()
	defer fake.uRLMutex.RUnlock()
	return len(fake.uRLArgsForCall)
}

func (fake *FakeS3Client) URLArgsForCall(i int) (string, string, bool, string) {
	fake.uRLMutex.RLock()
	defer fake.uRLMutex.RUnlock()
	return fake.uRLArgsForCall[i].bucketName, fake.uRLArgsForCall[i].remotePath, fake.uRLArgsForCall[i].private, fake.uRLArgsForCall[i].versionID
}

func (fake *FakeS3Client) URLReturns(result1 string) {
	fake.URLStub = nil
	fake.uRLReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeS3Client) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.bucketFilesMutex.RLock()
	defer fake.bucketFilesMutex.RUnlock()
	fake.bucketFileVersionsMutex.RLock()
	defer fake.bucketFileVersionsMutex.RUnlock()
	fake.uploadFileMutex.RLock()
	defer fake.uploadFileMutex.RUnlock()
	fake.downloadFileMutex.RLock()
	defer fake.downloadFileMutex.RUnlock()
	fake.deleteFileMutex.RLock()
	defer fake.deleteFileMutex.RUnlock()
	fake.deleteVersionedFileMutex.RLock()
	defer fake.deleteVersionedFileMutex.RUnlock()
	fake.uRLMutex.RLock()
	defer fake.uRLMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeS3Client) recordInvocation(key string, args []interface{}) {
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

var _ s3resource.S3Client = new(FakeS3Client)
