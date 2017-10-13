package blobstore

import (
	"encoding/json"
	"io/ioutil"
)

//go:generate counterfeiter -o fakes/fake_artifact.go . Artifact
type Artifact interface {
	Save(backup Backup) error
}

type FileArtifact struct {
	filePath string
}

func NewFileArtifact(filePath string) FileArtifact {
	return FileArtifact{filePath: filePath}
}

func (a FileArtifact) Save(backup Backup) error {
	marshalledBackup, err := json.MarshalIndent(backup, "", "  ")
	if err != nil {
		return err
	}

	ioutil.WriteFile(a.filePath, marshalledBackup, 0666)
	return nil
}

type Backup struct {
	DropletsBackup   BucketBackup `json:"droplets"`
	BuildpacksBackup BucketBackup `json:"buildpacks"`
	PackagesBackup   BucketBackup `json:"packages"`
}

type BucketBackup struct {
	BucketName string          `json:"bucket_name"`
	RegionName string          `json:"region_name"`
	Versions   []LatestVersion `json:"version"`
}

type LatestVersion struct {
	BlobKey string `json:"blob_key"`
	Id      string `json:"version_id"`
}
