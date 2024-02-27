package models

// import "gorm.io/gorm"

// type userBookmarks struct {
// 	User_id             int
// 	Bookmarked_tweet_id int
// }

// type userLikedTweets struct {
// 	User_id        int
// 	Liked_tweet_id int
// }

// func MigrateUserBookmarksNUserLikedTweets(db *gorm.DB) error {
// 	err1 := db.AutoMigrate(userBookmarks{})
// 	if err1 != nil {
// 		return err1
// 	}
// 	err2 := db.AutoMigrate(userLikedTweets{})
// 	if err2 != nil {
// 		return err2
// 	}
// 	return nil
// }
