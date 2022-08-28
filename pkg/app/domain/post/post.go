package postDomain

import (
	"fmt"
	"strings"

	"hexagonal_arch_with_Golang/pkg/app/api"
	"hexagonal_arch_with_Golang/pkg/config"
)

type Post struct {
	cfg        *config.Config
	fileName   string
	filePath   string
	fileUrl    string
	fileStatus string
	postStatus string
	postType   string
}

// Check if we actually implement relevant api
var _ api.Post = &Post{}

func New(cfg *config.Config, fileName, filePath, fileUrl, fileStatus, postStatus, postType string) *Post {
	return &Post{
		cfg:        cfg,
		fileName:   fileName,
		filePath:   filePath,
		fileUrl:    fileUrl,
		fileStatus: fileStatus,
		postStatus: postStatus,
		postType:   postType,
	}
}

func (ths *Post) FileParseUrl() {
	fileUrlParts := strings.Split(ths.fileUrl, "/")
	for _, part := range fileUrlParts {
		if strings.Contains(part, "jpg") {
			parts := strings.Split(part, ".")
			fmt.Printf("Name is: %s\n", parts[0])
			ths.fileName = parts[0]
			break
		}
	}
}

func (ths *Post) GetFileName() string {
	return ths.fileName
}

func (ths *Post) GetFileUrl() string {
	return ths.fileUrl
}

func (ths *Post) GetFileStatus() string {
	return ths.fileStatus
}

func (ths *Post) SetFileUrl(fileUrl string) {
	ths.fileUrl = fileUrl
}
