package routes

import (
	"back/auth"
	"back/model"
	"errors"
	"math"
	mrand "math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type attractionByIdResponse struct {
	Id uint 			`json:"id"`
	Name string 		`json:"name"`
	Description string 	`json:"description"`
	Photo string 		`json:"photo"`
	Rating float64 		`json:"rating"`
}

func (a *Api) AttractionById(c *gin.Context) {
	attractionIdByString := c.Param("attractionId")

	attractionId, err := strconv.Atoi(attractionIdByString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"Invalid request",
		})
		return
	}

	attraction, err := a.Db.Attraction.SelectById(uint(attractionId))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"message":"No attractions found with given id",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"message":"Service failure",
		})
		return
	}

	// calculate average rating 
	avg := averageRating(uint(attractionId), c, a)

	response := attractionByIdResponse{
		Id: attraction.Id,
		Name: attraction.Name,
		Description: attraction.Description,
		Photo: attraction.Photo,
		Rating: avg,
	}

	c.JSON(http.StatusOK, response)
}

type attractionWithRandomFunFactResponse struct {
	Id uint 		`json:"id"`
	Name string 	`json:"name"`
	FunFact string 	`json:"funfact"`
	Photo string 	`json:"photo"`
}

func (a *Api) AttractionWithRandomFunFact(c *gin.Context) {
	attractionsWithFunFact, err := a.Db.Attraction.SelectAllWithFunFact()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"message":"Service failure",
		})
		return
	}

	if len(attractionsWithFunFact) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message":"Not Found any attractions with fun fact",
		})
		return
	}

	randomAttractionNumber := generateRandomNumberInRange(0, uint(len(attractionsWithFunFact)-1))

	response := attractionWithRandomFunFactResponse{
		Id: attractionsWithFunFact[randomAttractionNumber].Id,
		Name: attractionsWithFunFact[randomAttractionNumber].Name,
		FunFact: attractionsWithFunFact[randomAttractionNumber].FunFact,
		Photo: attractionsWithFunFact[randomAttractionNumber].Photo,
	}
	
	c.JSON(http.StatusOK, response)
}

type createAttractionRequest struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Funfact string `json:"funfact"`
	Photo string `json:"photo"`
	LocationId uint `json:"locationId"`
	Type string `json:"type"`
}

type createAttractionResponse struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Funfact string `json:"funfact"`
	Photo string `json:"photo"`
	LocationId uint `json:"locationId"`
	Type string `json:"type"`
}

func (a *Api) CreateAttraction(c *gin.Context) {
	var request createAttractionRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"Invalid request",
		})
		return
	}

	// check if location type exists
	attractionType, err := a.Db.AttractionType.SelectByName(request.Type)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"message":"Type of attraction does not exist",
		})
		return
	}
	
	newAttraction := model.Attraction{
		Name: request.Name,
		Description: request.Description,
		FunFact: request.Funfact,
		Photo: request.Photo,
		RegionId: request.LocationId,
		AttractionTypeId: attractionType.Id,
		UserId: *auth.ContextId(c),
	}

	_, err = a.Db.Attraction.Create(newAttraction)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	c.JSON(http.StatusOK, newAttraction)
}

type attractionByLocationResponse struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Photo string `json:"photo"`
	Rating float32 `json:"rating"`
}

func (a *Api) AttractionByLocation(c *gin.Context) {
	locationIdAsString := c.Param("locationId")
	if locationIdAsString == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"Invalid request",
		})
		return
	}

	locationId, err := strconv.Atoi(locationIdAsString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"Invalid request",
		})
		return
	}

	category := c.Query("category")

	var attractions []model.Attraction
	if category == "" {
		attractions, err = a.Db.Attraction.SelectAllByLocationId(uint(locationId)) 
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"message":"Service failure",
			})
			return
		}
	} else {
		attractions, err = a.Db.Attraction.SelectAllByLocationIdAndCategory(uint(locationId), category) 
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"message":"Service failure",
			})
			return
		}
	}

	if len(attractions) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message":"No attractions found for given location",
		})
		return
	}
	
	var response []attractionByLocationResponse

	for _, item := range attractions {
		// calculate average rating
		avg := averageRating(item.Id, c, a)

		response = append(response, attractionByLocationResponse{
			Id: item.Id,
			Name: item.Name,
			Photo: item.Photo,
			Rating: float32(avg),
		})
	}

	c.JSON(http.StatusOK, response)
}

func generateRandomNumberInRange(min, max uint) uint {
    return uint(mrand.Intn(int(max-min+1)) + int(min))
}

func calculateAverageRating(ratings []model.Rating) float64{
	sum := 0.0;
	for _, item := range ratings {
		sum += float64(item.Rating)
	}
	avg := sum / float64(len(ratings))
	return math.Round(avg*100) / 100
}

func averageRating(attractionId uint, c *gin.Context, a *Api) float64{
	ratings, err := a.Db.Rating.SelectAllByAttractionId(attractionId)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"message":"Service failure",
		})
	}
	avg := calculateAverageRating(ratings)
	if len(ratings) == 0 {
		avg = 0.0
	}
	return avg
}