package api

type HelloWorld interface {
	HelloWorld(string, string) string
}

type Post interface {
	FileParseUrl()
	GetFileName() string
	GetFileUrl() string
	GetFileStatus() string
	SetFileUrl(fileUrl string)
}
