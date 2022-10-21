package controllers

import (
	"final_project/server/helper"
	"strconv"
	"time"

	"final_project/server/models"
	"final_project/server/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userPhoto struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type respPhoto struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	Photo_url string    `json:"photo_url"`
	UserId    int       `json:"user_id"`
	User      userPhoto
}

type PhotoController struct {
	repo repositories.PhotoRepo
}

func NewPhotoController(repo repositories.PhotoRepo) *PhotoController {
	return &PhotoController{repo: repo}

}

// Create Photo
// @Summary    Photo
// @Decription Photo
// @Tags       photo
// @Accept     json
// @Produce    json
// @Param Photo body models.ReqPhoto true "Photo"
// @Router     /photos [post]
// @Security BearerAuth
func (p *PhotoController) CreatePhoto(ctx *gin.Context) {
	getId := ctx.GetFloat64("id")
	var id int = int(getId)

	var req models.Photo
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	req.UserId = id
	errs := p.repo.CreatePhoto(&req)
	if errs != nil {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "ADD_PHOTO_FAIL",
			Error:   errs.Error(),
		})
		return
	}

	WriteJsonResponse(ctx, &helper.Response{
		Status:  http.StatusCreated,
		Message: "CREATE_PHOTO_SUCCESS",
		Payload: req,
	})
}

// Get Photo
// @Summary    Photo
// @Decription Photo
// @Tags       photo
// @Accept     json
// @Produce    json
// @Router     /photos [get]
// @Security BearerAuth
func (p *PhotoController) GetPhoto(ctx *gin.Context) {
	photos, err := p.repo.GetPhoto()
	if err != nil {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "GET_PHOTO_FAIL",
			Error:   err.Error(),
		})
	}
	var data []respPhoto
	for _, photo := range *photos {
		data = append(data, respPhoto{ID: photo.Id, Title: photo.Title, Caption: photo.Caption, Photo_url: photo.Photo_url, UserId: photo.UserId, CreatedAt: photo.CreatedAt, UpdatedAt: photo.UpdatedAt, User: userPhoto{Username: photo.Username, Email: photo.Email}})
	}
	WriteJsonResponse(ctx, &helper.Response{
		Status:  http.StatusOK,
		Message: "GET_PHOTO_SUCCESS",
		Payload: data,
	})
}

// Update Photo
// @Summary    Photo
// @Decription Photo
// @Tags       photo
// @Accept     json
// @Produce    json
// @Param Photo body models.ReqPhoto true "Photo"
// @Param      id path int true "Photo ID"
// @Router     /photos/{id} [put]
// @Security BearerAuth
func (p *PhotoController) UpdatePhoto(ctx *gin.Context) {
	getId := ctx.Params.ByName("photoid")
	id, errId := strconv.Atoi(getId)
	if errId != nil {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "UPDATE_PHOTO_FAIL",
			Error:   errId.Error(),
		})
		return
	}

	//check Author photo
	dataPhoto, errc := p.repo.GetPhotoById(id)
	if errc != nil {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "Photo tidak di temukan",
			Error:   errc.Error(),
		})
		return
	}

	getIdCtx := ctx.GetFloat64("id")
	var idUser int = int(getIdCtx)

	if dataPhoto.UserId != idUser {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "NO AUTHOR",
		})
		return
	}

	var req models.Photo
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	errUp := p.repo.UpdatePhoto(id, &req)
	if errUp != nil {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "UPDATE_PHOTO_FAIL",
			Error:   errUp.Error(),
		})
		return
	}

	WriteJsonResponse(ctx, &helper.Response{
		Status:  http.StatusOK,
		Message: "UPDATE_PHOTO_SUCCESS",
		Payload: req,
	})
}

// Delete Photo
// @Summary    Delete
// @Decription Delete
// @Tags       photo
// @Accept     json
// @Produce    json
// @Param      id path int true "Photo ID"
// @Router     /photos/{id} [delete]
// @Security BearerAuth
func (p *PhotoController) DeletePhoto(ctx *gin.Context) {
	getId := ctx.Params.ByName("photoid")
	id, errId := strconv.Atoi(getId)
	if errId != nil {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "DELETE_PHOTO_FAIL",
			Error:   errId.Error(),
		})
		return
	}

	//check Author photo
	dataPhoto, errc := p.repo.GetPhotoById(id)
	if errc != nil {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "Photo tidak di temukan",
			Error:   errc.Error(),
		})
		return
	}

	getIdCtx := ctx.GetFloat64("id")
	var idUser int = int(getIdCtx)

	if dataPhoto.UserId != idUser {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "NO AUTHOR",
		})
		return
	}

	err := p.repo.DeletePhoto(id)
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
		Message: "Your photo has been successfully deleted",
		Payload: id,
	})
}
