package models

type User struct{
	UserID   uint   `gorm:"primaryKey; autoIncrement"`
    Username string `gorm:"unique"`
    Password string
}