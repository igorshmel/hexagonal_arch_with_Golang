package domain

import (
	"hexagonal_arch_with_Golang/internal/domain/file"
	"hexagonal_arch_with_Golang/pkg/config"
)

type File interface {
	FileParseURL() string
	NewFileDomain(cfg *config.Config) *fileDomain.File
}
