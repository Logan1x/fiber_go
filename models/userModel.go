package models

type User struct {
	BaseModel
	Name     string
	Email    string    `gorm:"unique"`
	Projects []Project // Association with Project
}
