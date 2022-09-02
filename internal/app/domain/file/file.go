package fileDomain

import (
	"fmt"

	uuid2 "github.com/gofrs/uuid"
	"hexagonal_arch_with_Golang/pkg/config"
)

type File struct {
	cfg        *config.Config
	fileName   string
	filePath   string
	fileUrl    string
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

func (ths *File) FileParseUrl() {
	uuid, err := uuid2.NewV4()
	if err != nil {
		fmt.Printf("Fail generate uuidV4: %s\n", err.Error())
	}
	ths.fileName = fmt.Sprintf("%s%s", uuid.String(), ".jpg")

	/*	fileUrlParts := strings.Split(ths.fileUrl, "/")
		for _, part := range fileUrlParts {
			if strings.Contains(part, "jpg") {
				parts := strings.Split(part, ".")
				fmt.Printf("Name is: %s\n", parts[0])
				ths.fileName = parts[0]
				break
			}
		}
	*/
}

func (ths *File) GetFileName() string {
	return ths.fileName
}

func (ths *File) GetFilePath() string {
	return ths.filePath
}

func (ths *File) GetFileUrl() string {
	return ths.fileUrl
}

func (ths *File) GetFileHash() string {
	return ths.fileHash
}

func (ths *File) GetFileStatus() string {
	return ths.fileStatus
}

func (ths *File) SetFileName(fileName string) {
	ths.fileName = fileName
}

func (ths *File) SetFilePath(filePath string) {
	ths.filePath = filePath
}

func (ths *File) SetFileUrl(fileUrl string) {
	ths.fileUrl = fileUrl
}

func (ths *File) SetFileStatus(fileStatus string) {
	ths.fileStatus = fileStatus
}
