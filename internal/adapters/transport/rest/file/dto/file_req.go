package dto

import "github.com/gin-gonic/gin"

import validation "github.com/go-ozzo/ozzo-validation/v4"

// FileRequest is request
type FileRequest struct {
	FileUrl string `json:"file_url"`
}

// NewFileRequest is constructor
func NewFileRequest() *FileRequest {
	return &FileRequest{}
}

// Parse parses and validates the request
func (ths *FileRequest) Parse(c *gin.Context) error {
	return c.ShouldBindJSON(&ths)
}

// Validate validates an input request
func (ths *FileRequest) Validate() error {
	return validation.ValidateStruct(ths,
		validation.Field(&ths.FileUrl, validation.Required.Error("is required")),
	)
}
