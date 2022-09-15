package fileDomain

import (
	"fmt"

	"github.com/gofrs/uuid"
	"hexagonal_arch_with_Golang/pkg/config"
)

type File struct {
	cfg        *config.Config
	fileName   string
	filePath   string
	fileURL    string
	fileHash   string
	fileStatus string
}

// Check if we actually implement relevant api
//var _ application.File = &File{}

func New(cfg *config.Config) *File {
	return &File{
		cfg: cfg,
	}
}

func (ths *File) NewFileDomain(cfg *config.Config) *File {
	return &File{
		cfg: cfg,
	}
}

func (ths *File) FileParseURL() string {
	fileName, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Fail generate uuidV4: %s", err.Error())
	}

	return fmt.Sprintf("%s%s", fileName.String(), ".jpg")
}
