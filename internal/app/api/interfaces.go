package api

import (
	"hexagonal_arch_with_Golang/internal/app/domain/file"
	"hexagonal_arch_with_Golang/pkg/config"
)

type Notification interface {
	Notification(string, string) string
}

type File interface {
	FileParseUrl()
	GetFileName() string
	GetFilePath() string
	GetFileUrl() string
	GetFileHash() string
	GetFileStatus() string
	SetFileName(fileName string)
	SetFilePath(filePath string)
	SetFileUrl(fileUrl string)
	SetFileStatus(fileStatus string)
	NewFileDomain(cfg *config.Config) *fileDomain.File
}
