package routes

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type locationByIdResponse struct {
	Id uint 			`json:"id"`
	Name string			`json:"name"`
	Description string 	`json:"description"`
}

func (a *Api) LocationsById(c *gin.Context) {
	locationIdAsString := c.Param("locationId")
	if locationIdAsString == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid locationId",
		})
		return
	}

	locationId, err := strconv.Atoi(locationIdAsString)	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid locationId",
		})
		return
	}

	location, err := a.Db.Region.SelectById(uint(locationId))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Location not found",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"message": "Service failure",
		})
		return
	}

	response := locationByIdResponse{
		Id: location.Id,
		Name: location.Name,
		Description: location.Description,
	}
	c.JSON(http.StatusOK, response)
}

type loactionSearchResponse struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func (a *Api) LocationSearch(c *gin.Context) {
	searchQuery := c.Query("query")
	if searchQuery == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Parameter 'query' is required",
		})
		return
	}

	regions, err := a.Db.Region.SelectByNameFragment(searchQuery)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"message": "Service failure",
		})
		return
	}

	if len(regions) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
		return
	}

	var response []loactionSearchResponse
	for _, element := range regions {
		response = append(response, loactionSearchResponse{
			Id: element.Id,
			Name: element.Name,
			Type: element.RegionType.Name,
		})
	}

	c.JSON(http.StatusOK, response)
}
