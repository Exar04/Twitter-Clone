package initializers

import (
	"userService/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// type Repository struct {
// 	DB *gorm.DB
// }

func NewDbConnection() (*gorm.DB, error) {
	// dbHost := os.Getenv("DB_HOST")
	// dbPort := os.Getenv("DB_PORT")
	// dbUser := os.Getenv("DB_USER")
	// dbPassword := os.Getenv("DB_PASSWORD")
	// dbName := os.Getenv("DB_NAME")

	// connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	// 	dbHost, dbPort, dbUser, dbPassword, dbName)

	connStr := "host=localhost port=5432 user=yash password=yash dbname=twitterUserService sslmode=disable"
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	models.MigrateBlockedUsers(db)
	models.MigrateUserInfo(db)
	// models.MigrateUserBookmarksNUserLikedTweets(db)
	models.MigrateUserRelationship(db)

	return db, nil
}
