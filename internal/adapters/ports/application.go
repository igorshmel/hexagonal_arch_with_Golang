package ports

type AppPort interface {
	NewFile(fileUrl string) error
	Download(fileUrl, filePath, fileName string)
	Notification(name, message string)
}
