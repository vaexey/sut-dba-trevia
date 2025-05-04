package routes

import (
	"back/auth"
	"back/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createRatingRequest struct {
	AttractionId uint `json:"attractionId"`
	Rating int `json:"rating"`
}

type createRatingResponse struct {
	AttractionId uint `json:"attractionId"`
	UserId uint `json:"userId"`
	Rating int `json:"rating"`
}

func (a *Api) CreateRating(c *gin.Context) {
	var request createRatingRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"Invalid request",
		})
		return
	}

	if request.Rating < 1 || request.Rating > 5 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"Rating should be in range between 1 to 5",
		})
		return
	}

	// check if user already likes this attraction
	userRatings, err := a.Db.Rating.SelectByUserId(*auth.ContextId(c))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message":"User not found",
		})
		return
	}

	if isAttractionLikedByUser(request, userRatings) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"The user has already liked this attraction.",
		})
		return
	}

	_, err = a.Db.Rating.Create(model.Rating{
		Rating: request.Rating,
		UserId: *auth.ContextId(c),
		AttractionId: request.AttractionId,
	})
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"message":"Service failure",
		})
		return
	}

	response := createRatingResponse{
		AttractionId: request.AttractionId,
		UserId: *auth.ContextId(c),
		Rating: request.Rating,
	}

	c.JSON(http.StatusOK, response)
}

func isAttractionLikedByUser(request createRatingRequest, userRatings []model.Rating) bool {
	for _, item := range userRatings {
		if request.AttractionId == item.AttractionId {
			return true
		}
	}
	return false
}