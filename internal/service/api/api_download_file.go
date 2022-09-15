package api

import (
	"fmt"

	"hexagonal_arch_with_Golang/internal/adapters/stage/db/models"
	common "hexagonal_arch_with_Golang/pkg/common/file"
)

func (ths *Service) Download(fileURL, filePath, fileName string) {

	l := ths.cfg.Logger
	l.Info("start goroutine for Download:  %s", fileName)

	err := common.DownloadFile(fmt.Sprintf("%s%s", filePath, fileName), fileURL)
	if err != nil {
		l.Error("DownloadFile error  %s", err.Error())
	} else {
		l.Info("download success! %s", fileName)

		newFileModel := models.PsqlFile{
			FilePath:   filePath,
			FileName:   fileName,
			FileURL:    fileURL,
			FileStatus: "parse",
		}
		err = ths.db.NewRecordFile(&newFileModel)
		if err != nil {
			l.Info("DB Record fail: %s", err.Error())
		}
	}

	fileExist := ths.db.IsFileExists(fileName)
	if fileExist {
		err = ths.db.ChangeStatusFile(fileName, "download")
		if err != nil {
			l.Error("Error DB %s", err.Error())
		}
	}
}
