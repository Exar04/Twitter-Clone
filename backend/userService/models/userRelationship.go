package models

import (
	"gorm.io/gorm"
)

type UserRelationship struct {
	UserId      int `json:"userId"`
	FollowerId  int `json:"followerId"`
	FollowingId int `json:"followingId"`
}

func MigrateUserRelationship(db *gorm.DB) error {
	err := db.AutoMigrate(UserRelationship{})
	return err
}
