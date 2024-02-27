package models

import (
	"time"

	"gorm.io/gorm"
)

// when initially account is created it gets created by authService and this UserInfo user is added by gRPC
type UserInfo struct {
	User_id          int       `json:"id" gorm:"primary key; autoIncrement"`
	Username         string    `json:"username"`
	Email            string    `json:"email"`
	ProfilePicUrl    string    `json:"profilePic"`
	BackgroundPicUrl string    `json:"backgroundPicUrl"`
	Bio              string    `json:"bio"`
	Created_at       time.Time `json:"created_at"`
}

func MigrateUserInfo(db *gorm.DB) error {
	err := db.AutoMigrate(UserInfo{})
	return err
}
