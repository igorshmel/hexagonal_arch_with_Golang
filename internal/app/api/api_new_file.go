package api

import (
	"fmt"

	"hexagonal_arch_with_Golang/internal/models"
	"hexagonal_arch_with_Golang/pkg/dto/pb"
)

func (ths *Application) NewFile(fileUrl string) error {

	newFileDomain := ths.file.NewFileDomain(ths.cfg)

	newFileDomain.SetFileUrl(fileUrl)
	newFileDomain.SetFileStatus("parsing")
	newFileDomain.SetFilePath("../files/")

	newFileDomain.FileParseUrl()

	psqlFileModel := models.NewPsqlFile(
		newFileDomain.GetFileName(),
		newFileDomain.GetFilePath(),
		newFileDomain.GetFileUrl(),
		newFileDomain.GetFileHash(),
		newFileDomain.GetFileStatus(),
	)

	err := ths.db.NewRecordFile(psqlFileModel)
	if err != nil {
		ths.cfg.Logger.Error("Error DB %s\n", err.Error())
		fmt.Printf("Error DB %s\n", err.Error())
		return err
	}

	newFileProducer := pb.FileProducer{
		FileName:   newFileDomain.GetFileName(),
		FilePath:   newFileDomain.GetFilePath(),
		FileUrl:    fileUrl,
		FileStatus: newFileDomain.GetFileStatus(),
		Topic:      "file",
	}
	err = ths.kf.FileProducer(&newFileProducer)
	if err != nil {
		fmt.Printf("kafka FileProducer error  %s\n", err.Error())
		return err
	}

	newNotificationProducer := pb.NotificationProducer{
		Name:    newFileDomain.GetFileName(),
		Message: "New file",
		Topic:   "notification",
	}
	err = ths.kf.NotificationProducer(&newNotificationProducer)
	if err != nil {
		fmt.Printf("Error  %s\n", err.Error())
		return err
	}

	return nil
}
