package api

func (ths *Service) NewFile(fileURL string) error {

	fileDomain := ths.file.NewFileDomain(ths.cfg)
	fileName := fileDomain.FileParseURL()

	go ths.Download(
		fileURL,
		"../files/",
		fileName,
	)

	return nil
}
