package connection

import (
	"log"

	"animebase/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDB() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/animebase"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Studio{}, &models.Genre{}, &models.Anime{})

	return db
}
