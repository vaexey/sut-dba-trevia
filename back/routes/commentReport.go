package routes

import (
	"back/auth"
	"back/model"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type commentReportResponse struct {
	Id           uint
	UserId       uint
	CommentId    uint
	AttractionId uint
	Content      string
}

func (a *Api) CommentReports(c *gin.Context) {
	commentReports, err := a.Db.CommentReport.SelectAll()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"message": "Service failure"})
		return
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
		return
	}
	var response []commentReportResponse
	for _, report := range commentReports {
		// Fetch the comment to get AttractionId
		comment, err := a.Db.Comment.SelectById(report.CommentId)
		attractionId := uint(0)
		if err == nil && comment != nil {
			attractionId = comment.AttractionId
		}
		response = append(response, commentReportResponse{
			Id:           report.Id,
			UserId:       report.UserId,
			CommentId:    report.CommentId,
			AttractionId: attractionId,
			Content:      report.Content,
		})
	}
	c.JSON(http.StatusOK, response)
}

type createCommentReportRequest struct {
	CommentId uint
	Content   string
}

type createCommentReportResponse struct {
	UserId    uint
	CommentId uint
	Content   string
}

func (a *Api) CreateCommentReport(c *gin.Context) {
	var commentReportRequest createCommentReportRequest
	err := c.BindJSON(&commentReportRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	if commentReportRequest.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Content is empty."})
		return
	}

	// check if current user already created report for this comment
	currentUserReports, err := a.Db.CommentReport.SelectAllByUserIdAndCommentId(*auth.ContextId(c), commentReportRequest.CommentId)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"message": "Service failure"})
		return
	}

	if len(currentUserReports) != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Report already exists for this user"})
		return
	}

	commentReport := model.CommentReport{
		Content:   commentReportRequest.Content,
		UserId:    *auth.ContextId(c),
		CommentId: commentReportRequest.CommentId,
	}
	_, err = a.Db.CommentReport.Create(commentReport)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"message": "Service failure"})
		return
	}

	response := createCommentReportResponse{
		Content:   commentReport.Content,
		UserId:    *auth.ContextId(c),
		CommentId: commentReport.CommentId,
	}
	c.JSON(http.StatusOK, response)
}
