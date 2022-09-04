package file

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"hexagonal_arch_with_Golang/internal/adapters/transport/rest/file/dto"
)

// NewFileHandler is handler
func (ths Endpoint) NewFileHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		req := dto.NewFileRequest()

		// request parse
		if err := req.Parse(ctx); err != nil {
			ths.logger.Error("unable to parse a request: %s", err)
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		// validate request
		if err := req.Validate(); err != nil {
			ths.logger.Error("error of validation: %s", err)
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err := ths.app.NewFile(req.FileUrl)
		if err != nil {
			ths.logger.Error("filed Call toi app.NewFile(): %s", err)
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		resp := dto.NewFileResponse()

		ctx.JSON(http.StatusOK, gin.H{"file": resp})
	}
}
