package db

import (
	"fmt"
	"os"

	"final_project/server/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectGorm() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}
	host := os.Getenv("HOST")
	port := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	pass := os.Getenv("PASS")
	dbname := os.Getenv("DBNAME")
	fmt.Println(host, port, user, pass, dbname)

	confPg := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, dbname,
	)

	db, errDb := gorm.Open(postgres.Open(confPg))
	if errDb != nil {
		panic(errDb)
	}
	db.Debug().AutoMigrate(models.User{})
	db.Debug().AutoMigrate(models.Photo{})
	db.Debug().AutoMigrate(models.SocialMedia{})
	db.Debug().AutoMigrate(models.Comment{})

	return db

}
