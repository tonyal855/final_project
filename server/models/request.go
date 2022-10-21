package models

type ReqRegister struct {
	Username string `json:"username"  binding:"required"`
	Email    string `json:"email"  binding:"required,email"`
	Password string `json:"password"  binding:"required,min=6"`
	Age      int    `json:"age"  binding:"required,min=9"`
}

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ReqUserUpdate struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type ReqPhoto struct {
	Title     string `json:"title" binding:"required"`
	Caption   string `json:"caption"`
	Photo_url string `json:"photo_url"  binding:"required"`
}

type ReqSocmed struct {
	Name             string `json:"name" binding:"required"`
	Social_media_url string `json:"social_media_url" binding:"required"`
}

type ReqComment struct {
	Message string `json:"message" binding:"required"`
	PhotoId int    `json:"photo_id"`
}
