package application

import (
	"fmt"

	common "hexagonal_arch_with_Golang/pkg/adapters/common/file"
	"hexagonal_arch_with_Golang/pkg/config"
	"hexagonal_arch_with_Golang/pkg/dto/pb"
	"hexagonal_arch_with_Golang/pkg/ports"
)

type Application struct {
	cfg          *config.Config
	db           ports.DbPort
	kf           ports.KafkaPort
	notification Notification
	file         File
}

// Check if we actually implement all the ports.
//var _ ports.AppPort = &Application{}

func New(cfg *config.Config, db ports.DbPort, kf ports.KafkaPort, notification Notification, file File) *Application {
	return &Application{
		cfg:          cfg,
		db:           db,
		kf:           kf,
		notification: notification,
		file:         file,
	}
}

func (ths *Application) AppFile(fileUrl string) error {
	newFileDomain := ths.file.NewFileDomain(ths.cfg)
	newFileDomain.SetFileUrl(fileUrl)
	newFileDomain.SetFileStatus("parsing")
	newFileDomain.SetFilePath("../files/")
	newFileDomain.FileParseUrl()

	psqlFileModel := ths.db.NewPsqlFile(
		newFileDomain.GetFileName(),
		newFileDomain.GetFilePath(),
		newFileDomain.GetFileUrl(),
		newFileDomain.GetFileHash(),
		newFileDomain.GetFileStatus(),
	)

	err := ths.db.NewRecordFile(psqlFileModel)
	if err != nil {
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
		fmt.Printf("Error  %s\n", err.Error())
		return err
	}
	return nil
}

func (ths *Application) Download(fileUrl, filePath, fileName string) {

	fmt.Printf("\nstart gorutine for Download:  %s\n", fileName)

	err := common.DownloadFile(fmt.Sprintf("%s%s", filePath, fileName), fileUrl)
	if err != nil {
		fmt.Printf("Error  %s\n", err.Error())
		err = ths.kf.NotificationProducer(fmt.Sprintf("download fail! %s", fileName), "notification")
		if err != nil {
			fmt.Printf("Error  %s\n", err.Error())
		}
	} else {
		fmt.Printf("\ndownload success! %s\n", fileName)
		err = ths.kf.NotificationProducer(fmt.Sprintf("download success! %s", fileName), "notification")
		if err != nil {
			fmt.Printf("Error  %s\n", err.Error())
		}
	}
}

func (ths *Application) Notification(name, message string) {

	psqlNotificationModel := ths.db.NewPsqlNotification(
		name,
		message,
	)
	err := ths.db.NotificationRecord(psqlNotificationModel)
	if err != nil {
		fmt.Printf("Error DB %s\n", err.Error())
	}

}
