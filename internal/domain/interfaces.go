package domain

import (
	"hexagonal_arch_with_Golang/internal/domain/file"
	"hexagonal_arch_with_Golang/pkg/config"
)

type Notification interface {
	Notification(string, string) string
	GetName() string
	GetMassage() string
	SetMassage(string)
	SetName(string)
}

type File interface {
	FileParseURL()
	GetFileName() string
	GetFilePath() string
	GetFileURL() string
	GetFileHash() string
	GetFileStatus() string
	SetFileName(fileName string)
	SetFilePath(filePath string)
	SetFileURL(fileURL string)
	SetFileStatus(fileStatus string)
	NewFileDomain(cfg *config.Config) *fileDomain.File
}
