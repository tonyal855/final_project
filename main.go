package main

import (
	"final_project/db"
	"final_project/server"
	"final_project/server/controllers"
	"final_project/server/repositories/gorm"
	"os"
)

// @description    API MyGram
// @termsOfService http://swagger.io/terms/
// @BasePath       /
// @contact.name   Toni
// @contact.email  toni.al855@gmail.com
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	db := db.ConnectGorm()

	photoRepo := gorm.NewPhotoRepo(db)
	photocontroller := controllers.NewPhotoController(photoRepo)

	socialMediaRepo := gorm.NewSocialMediaRepo(db)
	socialMediaController := controllers.NewSocialMediaController(socialMediaRepo, photoRepo)

	commentRepo := gorm.NewCommentRepo(db)
	commentController := controllers.NewCommentController(commentRepo)

	userRepo := gorm.NewUserRepo(db)
	userController := controllers.NewUserController(userRepo, photoRepo, commentRepo, socialMediaRepo)
	router := server.NewRouter(userController, photocontroller, socialMediaController, commentController)
	var PORT = os.Getenv("PORTP")
	router.Start(":" + PORT)

}
