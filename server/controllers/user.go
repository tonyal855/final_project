package controllers

import (
	"final_project/server/helper"
	"final_project/server/models"
	"final_project/server/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	repo        repositories.UserRepo
	repoPhoto   repositories.PhotoRepo
	repoComment repositories.CommentRepo
	repoSocmed  repositories.SocialMediaRepo
}

func NewUserController(repo repositories.UserRepo, repoPhoto repositories.PhotoRepo, repoComment repositories.CommentRepo, repoSocmed repositories.SocialMediaRepo) *UserController {
	return &UserController{repo: repo,
		repoPhoto:   repoPhoto,
		repoComment: repoComment,
		repoSocmed:  repoSocmed,
	}
}

func WriteJsonResponse(ctx *gin.Context, payload *helper.Response) {
	ctx.JSON(payload.Status, payload)
}

// Register
// @Summary    Register
// @Decription Register
// @Tags       users
// @Accept     json
// @Produce    json
// @Param Register body models.ReqRegister true "Register"
// @Router     /users/register [post]
func (u *UserController) CreateUser(ctx *gin.Context) {
	var req models.User
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//check email and username
	_, errU := u.repo.GetUserByUser(req.Username)
	if errU == nil {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "username sudah ada",
		})
		return
	}

	_, errE := u.repo.GetUserByEmail(req.Email)
	if errE == nil {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "email sudah ada",
		})
		return
	}

	req.Password = helper.GeneratePasswordBrypt(req.Password)
	errs := u.repo.CreateUser(&req)
	if errs != nil {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "ADD_USER_FAIL",
			Error:   errs.Error(),
		})
		return
	}
	var resp = models.ReqUser{ID: int(req.ID), Username: req.Username, Email: req.Email, Age: req.Age}
	WriteJsonResponse(ctx, &helper.Response{
		Status:  http.StatusCreated,
		Message: "CREATE_USER_SUCCESS",
		Payload: resp,
	})
}

// Login
// @Summary    Login
// @Decription Login
// @Tags       users
// @Accept     json
// @Produce    json
// @Param login body models.ReqLogin true "Login"
// @Router     /users/login [post]
func (u *UserController) Login(ctx *gin.Context) {
	var req models.ReqLogin
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, errU := u.repo.GetUserByEmail(req.Email)
	if errU != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Email Salah",
		})
		return
	}

	comparePass := helper.ComparePwd([]byte(user.Password), []byte(req.Password))
	if !comparePass {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Password Salah",
		})
		return
	}
	token, errT := helper.GenerateToken(user.ID, user.Email)
	if errT != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": errT.Error(),
		})
		return
	}
	WriteJsonResponse(ctx, &helper.Response{
		Status: http.StatusOK,
		Token:  token,
	})
}

// Update User
// @Summary    Update User
// @Decription Update
// @Tags       users
// @Accept     json
// @Produce    json
// @Param Update body models.ReqUserUpdate true "Update User"
// @Router     /users [put]
// @Security BearerAuth
func (u *UserController) UpdateUser(ctx *gin.Context) {

	getId := ctx.GetFloat64("id")
	var id int = int(getId)

	var req models.ReqUser
	er := ctx.ShouldBindJSON(&req)
	if er != nil {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "UPDATE_PEOPLE_FAIL",
			Error:   er.Error(),
		})
		return
	}

	//check email and username
	_, errU := u.repo.GetUserByUser(req.Username)
	if errU == nil {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "username sudah ada",
		})
		return
	}

	_, errE := u.repo.GetUserByEmail(req.Email)
	if errE == nil {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "email sudah ada",
		})
		return
	}

	errUp := u.repo.UpdateUser(id, &req)

	if errUp != nil {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "UPDATE_USER_FAIL",
			Error:   errUp.Error(),
		})
		return
	}

	WriteJsonResponse(ctx, &helper.Response{
		Status:  http.StatusOK,
		Message: "UPDATE_USER_SUCCESS",
		Payload: req,
	})
}

// Delete User
// @Summary    Delete User
// @Decription Delete
// @Tags       users
// @Accept     json
// @Produce    json
// @Router     /users [delete]
// @Security BearerAuth
func (u *UserController) DeleteUser(ctx *gin.Context) {
	getId := ctx.GetFloat64("id")
	var id int = int(getId)

	errP := u.repoPhoto.DeletePhotoByUserId(id)
	if errP != nil {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "DELETE_USER_FAIL",
			Error:   errP.Error(),
		})
		return
	}

	errC := u.repoComment.DeleteCommentByUserId(id)
	if errC != nil {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "DELETE_USER_FAIL",
			Error:   errC.Error(),
		})
		return
	}

	errS := u.repoSocmed.DeleteSocmedByUserId(id)
	if errS != nil {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "DELETE_USER_FAIL",
			Error:   errS.Error(),
		})
		return
	}

	err := u.repo.DeleteUser(id)
	if err != nil {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "DELETE_USER_FAIL",
			Error:   err.Error(),
		})
		return
	}
	WriteJsonResponse(ctx, &helper.Response{
		Status:  http.StatusOK,
		Message: "DELETE_USER_SUCCESS",
		Payload: id,
	})
}
