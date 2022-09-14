package api

import (
	"fmt"

	common "hexagonal_arch_with_Golang/pkg/common/file"
	"hexagonal_arch_with_Golang/pkg/dto/pb"
)

func (ths *Service) Download(fileURL, filePath, fileName string) {
	l := ths.cfg.Logger
	l.Info("start goroutine for Download:  %s", fileName)

	err := common.DownloadFile(fmt.Sprintf("%s%s", filePath, fileName), fileURL)
	if err != nil {
		l.Error("DownloadFile error  %s", err.Error())

		newNotificationProducer := pb.NotificationProducer{
			Name:    fileName,
			Message: "error download file",
			Topic:   "notification",
		}
		err = ths.kf.NotificationProducer(&newNotificationProducer)
		if err != nil {
			l.Error("Error  %s", err.Error())
		}
	} else {
		l.Info("download success! %s", fileName)

		newNotificationProducer := pb.NotificationProducer{
			Name:    fileName,
			Message: "success download file",
			Topic:   "notification",
		}
		err = ths.kf.NotificationProducer(&newNotificationProducer)
		if err != nil {
			l.Error("Error  %s", err.Error())
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
