package controllers

import (
	"final_project/server/helper"
	"final_project/server/models"
	"final_project/server/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type userComment struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type photoComment struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Caption   string `json:"Caption"`
	Photo_url string `json:"photo_url"`
	User_id   int    `json:"user_id"`
}

type respComment struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Message   string    `json:"message"`
	Photo_id  int       `json:"photo_id"`
	User_id   int       `json:"user_id"`
	User      userComment
	Photo     photoComment
}

type CommentController struct {
	repo repositories.CommentRepo
}

func NewCommentController(repo repositories.CommentRepo) *CommentController {
	return &CommentController{repo: repo}

}

// Create Comment
// @Summary    Comment
// @Decription Comment
// @Tags       Comment
// @Accept     json
// @Produce    json
// @Param Comment body models.ReqComment true "Comment"
// @Router     /comments [post]
// @Security BearerAuth
func (c *CommentController) CreateComment(ctx *gin.Context) {
	getId := ctx.GetFloat64("id")
	var id int = int(getId)

	var req models.Comment
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	req.UserId = id
	errs := c.repo.CreateComment(&req)
	if errs != nil {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "ADD_COMMENT_FAIL",
			Error:   errs.Error(),
		})
		return
	}

	WriteJsonResponse(ctx, &helper.Response{
		Status:  http.StatusCreated,
		Message: "CREATE_COMMENT_SUCCESS",
		Payload: req,
	})
}

// Get Comment
// @Summary    Comment
// @Decription Comment
// @Tags       Comment
// @Accept     json
// @Produce    json
// @Router     /comments [get]
// @Security BearerAuth
func (c *CommentController) GetComment(ctx *gin.Context) {
	comments, err := c.repo.GetComment()
	if err != nil {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "GET_COMMENT_FAIL",
			Error:   err.Error(),
		})
	}
	var data []respComment
	for _, comment := range *comments {
		data = append(data, respComment{Id: comment.Id, CreatedAt: comment.CreatedAt, UpdatedAt: comment.UpdatedAt, Message: comment.Message, Photo_id: comment.PhotoId, User_id: comment.UserId, User: userComment{Id: comment.User_id_user, Email: comment.Email, Username: comment.Username}, Photo: photoComment{Id: comment.Photo_id_photo, Title: comment.Title, Caption: comment.Caption, Photo_url: comment.Photo_url, User_id: comment.User_id_photo}})
	}
	WriteJsonResponse(ctx, &helper.Response{
		Status:  http.StatusOK,
		Message: "GET_COMMENT_SUCCESS",
		Payload: data,
	})
}

// Update Comment
// @Summary    Comment
// @Decription Comment
// @Tags       Comment
// @Accept     json
// @Produce    json
// @Param Comment body models.ReqComment true "Comment"
// @Param      id path int true "Comment ID"
// @Router     /comments/{id} [put]
// @Security BearerAuth
func (c *CommentController) UpdateComment(ctx *gin.Context) {
	getId := ctx.Params.ByName("commentId")
	id, errId := strconv.Atoi(getId)
	if errId != nil {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "UPDATE_COMMENT_FAIL",
			Error:   errId.Error(),
		})
		return
	}

	//check Author photo
	dataComment, errc := c.repo.GetCommentById(id)
	if errc != nil {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "Comment tidak di temukan",
			Error:   errc.Error(),
		})
		return
	}

	getIdCtx := ctx.GetFloat64("id")
	var idUser int = int(getIdCtx)

	if dataComment.UserId != idUser {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "NO AUTHOR",
		})
		return
	}

	var req models.Comment
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	errUp := c.repo.UpdateComment(id, &req)
	if errUp != nil {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "UPDATE_COMMENT_FAIL",
			Error:   errUp.Error(),
		})
		return
	}

	WriteJsonResponse(ctx, &helper.Response{
		Status:  http.StatusOK,
		Message: "UPDATE_COMMENT_SUCCESS",
		Payload: req,
	})
}

// Delete Comment
// @Summary    Comment
// @Decription Comment
// @Tags       Comment
// @Accept     json
// @Produce    json
// @Param      id path int true "Comment ID"
// @Router     /comments/{id} [delete]
// @Security BearerAuth
func (c *CommentController) DeleteComment(ctx *gin.Context) {
	getId := ctx.Params.ByName("commentId")
	id, errId := strconv.Atoi(getId)
	if errId != nil {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "DELETE_COMMENT_FAIL",
			Error:   errId.Error(),
		})
		return
	}

	//check Author photo
	dataComment, errc := c.repo.GetCommentById(id)
	if errc != nil {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "Comment tidak di temukan",
			Error:   errc.Error(),
		})
		return
	}

	getIdCtx := ctx.GetFloat64("id")
	var idUser int = int(getIdCtx)

	if dataComment.UserId != idUser {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "NO AUTHOR",
		})
		return
	}

	err := c.repo.DeleteComment(id)
	if err != nil {
		WriteJsonResponse(ctx, &helper.Response{
			Status:  http.StatusInternalServerError,
			Message: "DELETE_COMMENT_FAIL",
			Error:   err.Error(),
		})
		return
	}
	WriteJsonResponse(ctx, &helper.Response{
		Status:  http.StatusOK,
		Message: "Your Comment has been successfully deleted",
		Payload: id,
	})
}
