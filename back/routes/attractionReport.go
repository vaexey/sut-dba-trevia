package routes

import (
	"back/auth"
	"back/model"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type attractionReportResponse struct {
	Id uint
	Content string
	AttractionId uint
	UserId uint
}

func (a *Api) AttractionReports(c *gin.Context) {
	attractionReports, err := a.Db.AttractionReport.SelectAll()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"message":"Service failure"})
		return
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"message":"Not found"})
		return
	}

	var response []attractionReportResponse
	for _, report := range attractionReports {
		response = append(response, attractionReportResponse{
			Id: report.Id,
			Content: report.Content,
			AttractionId: report.AttractionId,
			UserId: report.UserId,
		})
	}
	c.JSON(http.StatusOK, response)
}

type createAttractionReportRequest struct {
	AttractionId uint
	Content string
}

type createAttractionReportResponse struct {
	Content string
	AttractionId uint
	UserId uint
}

func (a *Api) CreateAttractionReport(c *gin.Context) {
	var attractionReportRequest createAttractionReportRequest
	err := c.BindJSON(&attractionReportRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message":"Bad request"})
		return
	}
	if attractionReportRequest.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message":"Content is empty."})
		return
	}

	// check if current user already created report for this attraction
	currentUserReports, err := a.Db.AttractionReport.SelectAllByUserIdAndAttractionId(*auth.ContextId(c), attractionReportRequest.AttractionId)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"message":"Service failure"})
		return
	}

	if len(currentUserReports) != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message":"Report already exists for this user"})
		return
	}

	attractionReport := model.AttractionReport{
		Content: attractionReportRequest.Content,
		AttractionId: attractionReportRequest.AttractionId,
		UserId: *auth.ContextId(c),
	}

	_, err = a.Db.AttractionReport.Create(attractionReport)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"message":"Service failure"})
		return
	}

	response := createAttractionReportResponse{
		Content: attractionReport.Content,
		AttractionId: attractionReport.AttractionId,
		UserId: *auth.ContextId(c),
	}
	c.JSON(http.StatusOK, response)
}