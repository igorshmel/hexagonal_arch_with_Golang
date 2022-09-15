package service

type ApiPort interface {
	NewFile(fileURL string) error
	Download(fileURL, filePath, fileName string)
}
