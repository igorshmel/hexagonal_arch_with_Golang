package api

import (
	"hexagonal_arch_with_Golang/pkg/dto/pb"
)

func (ths *Application) NewFile(fileUrl string) error {
	l := ths.cfg.Logger

	fileDomain := ths.file.NewFileDomain(ths.cfg)

	fileDomain.SetFileUrl(fileUrl)
	fileDomain.SetFileStatus("parsing")
	fileDomain.SetFilePath("../files/")

	fileDomain.FileParseUrl()

	err := ths.db.NewRecordFile(fileDomain)
	if err != nil {
		l.Error("Error DB %s", err.Error())
		return err
	}

	fileProducer := pb.FileProducer{
		FileName:   fileDomain.GetFileName(),
		FilePath:   fileDomain.GetFilePath(),
		FileUrl:    fileUrl,
		FileStatus: fileDomain.GetFileStatus(),
		Topic:      "file",
	}
	err = ths.kf.FileProducer(&fileProducer)
	if err != nil {
		l.Error("kafka FileProducer error  %s", err.Error())
		return err
	}

	notificationProducer := pb.NotificationProducer{
		Name:    fileDomain.GetFileName(),
		Message: "New file",
		Topic:   "notification",
	}
	err = ths.kf.NotificationProducer(&notificationProducer)
	if err != nil {
		l.Error("Error  %s", err.Error())
		return err
	}

	return nil
}
