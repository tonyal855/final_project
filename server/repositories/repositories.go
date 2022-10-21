package repositories

import "final_project/server/models"

type UserRepo interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByUser(username string) (*models.User, error)
	UpdateUser(id int, userReq *models.ReqUser) error
	DeleteUser(id int) error
}

type PhotoRepo interface {
	CreatePhoto(photo *models.Photo) error
	GetPhoto() (*[]models.Photos, error)
	UpdatePhoto(id int, req *models.Photo) error
	GetPhotoById(id int) (*models.Photo, error)
	DeletePhoto(id int) error
	GetPhotoByUserId(id int) (*models.Photo, error)
	DeletePhotoByUserId(user_id int) error
}

type SocialMediaRepo interface {
	CreateSocialMedia(photo *models.SocialMedia) error
	GetSocialMedia() (*[]models.SocialMedias, error)
	UpdateSocialMedia(id int, req *models.SocialMedia) error
	DeleteSocMed(id int) error
	GetSocMedById(id int) (*models.SocialMedia, error)
	DeleteSocmedByUserId(user_id int) error
}

type CommentRepo interface {
	CreateComment(comment *models.Comment) error
	GetComment() (*[]models.Comments, error)
	UpdateComment(id int, req *models.Comment) error
	GetCommentById(id int) (*models.Comment, error)
	DeleteComment(id int) error
	DeleteCommentByUserId(user_id int) error
}
