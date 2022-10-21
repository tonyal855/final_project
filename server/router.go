package server

import (
	"final_project/docs"
	"final_project/server/controllers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	userRouter        *controllers.UserController
	photoRouter       *controllers.PhotoController
	socialMediaRouter *controllers.SocialMediaController
	commentRouter     *controllers.CommentController
}

func NewRouter(useRouter *controllers.UserController, photoRouter *controllers.PhotoController, socialMediaRouter *controllers.SocialMediaController, commentRouter *controllers.CommentController) *Router {
	return &Router{
		userRouter:        useRouter,
		photoRouter:       photoRouter,
		socialMediaRouter: socialMediaRouter,
		commentRouter:     commentRouter,
	}
}

func (r *Router) Start(port string) {
	docs.SwaggerInfo.Title = "Swagger MyGram API"
	docs.SwaggerInfo.Description = "Sample API Spec for MyGram"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "https://finalproject-production-9000.up.railway.app"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.Default()
	router.POST("/users/register", r.userRouter.CreateUser)
	router.POST("/users/login", r.userRouter.Login)
	router.PUT("/users", CheckAuth, r.userRouter.UpdateUser)
	router.DELETE("/users", CheckAuth, r.userRouter.DeleteUser)

	router.POST("/photos", CheckAuth, r.photoRouter.CreatePhoto)
	router.GET("/photos", CheckAuth, r.photoRouter.GetPhoto)
	router.PUT("/photos/:photoid", CheckAuth, r.photoRouter.UpdatePhoto)
	router.DELETE("/photos/:photoid", CheckAuth, r.photoRouter.DeletePhoto)

	router.POST("/socialmedias", CheckAuth, r.socialMediaRouter.CreateSocialMedia)
	router.GET("/socialmedias", CheckAuth, r.socialMediaRouter.GetSocialMedia)
	router.PUT("/socialmedias/:socialMediaId", CheckAuth, r.socialMediaRouter.UpdateSocialMedia)
	router.DELETE("/socialmedias/:socialMediaId", CheckAuth, r.socialMediaRouter.DeleteSocmed)

	router.POST("/comments", CheckAuth, r.commentRouter.CreateComment)
	router.GET("/comments", CheckAuth, r.commentRouter.GetComment)
	router.PUT("/comments/:commentId", CheckAuth, r.commentRouter.UpdateComment)
	router.DELETE("/comments/:commentId", CheckAuth, r.commentRouter.DeleteComment)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(port)
}
