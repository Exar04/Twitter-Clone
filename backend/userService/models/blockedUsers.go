package models

import "gorm.io/gorm"

type BlockedUsers struct {
	User_id        int `json:"id"`
	BlockedUser_id int `json:"blockedUserId"`
}

func MigrateBlockedUsers(db *gorm.DB) error {
	err := db.AutoMigrate(BlockedUsers{})
	return err
}
