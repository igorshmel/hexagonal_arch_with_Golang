package app

type ApiPort interface {
	NewFile(fileURL string) error
	Download(fileURL, filePath, fileName string)
	Notification(name, message string)
}
