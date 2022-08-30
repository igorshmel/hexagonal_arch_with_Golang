package models

const filesTableName = "files"
const notificationTableName = "notification"

type PsqlNotification struct {
	BaseModel
	Name    string `gorm:"type:varchar(255)"`
	Message string `gorm:"type:varchar(255)"`
}

type PsqlFile struct {
	BaseModel
	FileName   string `gorm:"type:varchar(255)"`
	FilePath   string `gorm:"type:text"`
	FileUrl    string `gorm:"type:text"`
	FileHash   string `gorm:"type:text"`
	FileStatus string `gorm:"type:varchar(255)"`
}

// TableName returns the name of the table
func (PsqlFile) TableName() string {
	return filesTableName
}

// TableName returns the name of the table
func (PsqlNotification) TableName() string {
	return notificationTableName
}
