package models

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex"`
	Email    string `gorm:"uniqueIndex"`
	Password string
}

var db *gorm.DB

func InitDB(dsn string) error {
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return db.AutoMigrate(&User{})
}

func CreateUser(user User) error {
	result := db.Create(&user)
	return result.Error
}

func GetUserByUsername(username string) (User, error) {
	var user User
	result := db.Where("username = ?", username).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return User{}, errors.New("user not found")
	}
	return user, result.Error
}

func GetUserByEmail(email string) (User, error) {
	var user User
	result := db.Where("email = ?", email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return User{}, errors.New("user not found")
	}
	return user, result.Error
}
