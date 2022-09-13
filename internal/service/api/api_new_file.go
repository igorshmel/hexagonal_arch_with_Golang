package api

import (
	"hexagonal_arch_with_Golang/pkg/dto/pb"
)

func (ths *Service) NewFile(fileURL string) error {
	l := ths.cfg.Logger

	fileDomain := ths.file.NewFileDomain(ths.cfg)

	fileDomain.SetFileURL(fileURL)
	fileDomain.SetFileStatus("parsing")
	fileDomain.SetFilePath("../files/")

	fileDomain.FileParseURL()

	err := ths.db.NewRecordFile(fileDomain)
	if err != nil {
		l.Error("Error DB %s", err.Error())
		return err
	}

	fileProducer := pb.FileProducer{
		FileName:   fileDomain.GetFileName(),
		FilePath:   fileDomain.GetFilePath(),
		FileUrl:    fileURL,
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
