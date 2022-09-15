package ports

type RestPort interface {
	Run()
	FileHandlers()
}
