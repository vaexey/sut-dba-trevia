package routes

import (
	"back/auth"
	"back/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createCommentRequest struct {
	AttractionId uint   `json:"attractionId" binding:"required"`
	Comment      string `json:"comment"      binding:"required"`
}

type commentResponse struct {
	AttractionId uint   `json:"attractionId"`
	UserId       uint   `json:"userId"`
	Username     string `json:"username"`
	Comment      string `json:"comment"`
}

func (a *Api) GetComments(c *gin.Context) {
	idParam := c.Param("attractionId")
	attrID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid attraction id"})
		return
	}

	comments, err := a.Db.Comment.SelectAllByAttractionId(uint(attrID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "service failure"})
		return
	}

	var resp []commentResponse
	for _, cm := range comments {
		resp = append(resp, commentResponse{
			AttractionId: cm.AttractionId,
			UserId:       cm.UserId,
			Username:     cm.User.Username,
			Comment:      cm.Comment,
		})
	}
	c.JSON(http.StatusOK, resp)
}

func (a *Api) CreateComment(c *gin.Context) {
	var req createCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	uid := auth.ContextId(c)
	if uid == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	newComment := model.Comment{
		AttractionId: req.AttractionId,
		UserId:       *uid,
		Comment:      req.Comment,
	}

	_, err := a.Db.Comment.Create(newComment)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"message": "service failure"})
		return
	}

	user, err := a.Db.User.SelectById(*uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not load user"})
		return
	}

	resp := commentResponse{
		AttractionId: req.AttractionId,
		UserId:       *uid,
		Username:     user.Username,
		Comment:      req.Comment,
	}
	c.JSON(http.StatusOK, resp)
}
