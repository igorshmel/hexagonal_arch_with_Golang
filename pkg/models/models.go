package models

type Greeting struct {
	BaseModel
	Greeting string `gorm:"type:varchar(255)"`
}

type File struct {
	BaseModel
	FileName string `gorm:"type:varchar(255)"`
	FileUrl  string `gorm:"type:varchar(255)"`
}
