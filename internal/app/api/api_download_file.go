package api

import (
	"fmt"

	common "hexagonal_arch_with_Golang/pkg/common/file"
	"hexagonal_arch_with_Golang/pkg/dto/pb"
)

func (ths *Application) Download(fileUrl, filePath, fileName string) {

	fmt.Printf("\nstart gorutine for Download:  %s\n", fileName)

	err := common.DownloadFile(fmt.Sprintf("%s%s", filePath, fileName), fileUrl)
	if err != nil {
		fmt.Printf("DownladFile error  %s\n", err.Error())

		newNotificationProducer := pb.NotificationProducer{
			Name:    fileName,
			Message: "error download file",
			Topic:   "notification",
		}
		err = ths.kf.NotificationProducer(&newNotificationProducer)
		if err != nil {
			fmt.Printf("Error  %s\n", err.Error())
		}
	} else {
		fmt.Printf("\ndownload success! %s\n", fileName)

		newNotificationProducer := pb.NotificationProducer{
			Name:    fileName,
			Message: "success download file",
			Topic:   "notification",
		}
		err = ths.kf.NotificationProducer(&newNotificationProducer)
		if err != nil {
			fmt.Printf("Error  %s\n", err.Error())
		}
	}

	fileExist := ths.db.IsFileExists(fileName)
	if fileExist {
		err = ths.db.ChangeStatusFile(fileName, "download")
		if err != nil {
			fmt.Printf("Error DB %s\n", err.Error())
		}
	}
}
