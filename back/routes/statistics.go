package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
)

type statsAttractionResponse struct {
	Name             string
	NumberOfComments uint
}

type statsUserResponse struct {
	Name             string
	NumberOfComments uint
}

func (a *Api) Stats(c *gin.Context) {
	attractionRecords, userRecords := recordsNumber(c)

	attractions, err := a.Db.Attraction.SelectAttractionsWithMostComments(attractionRecords)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"message": "Service failure"})
		return
	}
	users, err := a.Db.User.SelectUsersWithMostComments(userRecords)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"message": "Service failure"})
		return
	}

	// create pdf
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 14)

	// attractions
	pdf.Cell(40, 10, "Top 5 Attractions by Number of Comments")
	pdf.Ln(12)
	pdf.SetFont("Arial", "", 12)
	for _, item := range attractions {
		count, err := a.Db.Comment.CountCommentsByAttractionId(item.Id)
		if err != nil {
			continue // skip wrong data
		}
		line := item.Name + " - comments: " + strconv.FormatUint(uint64(count), 10)
		pdf.Cell(0, 10, line)
		pdf.Ln(8)
	}

	// users
	pdf.Ln(10)
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(40, 10, "Top 5 Users by Number of Comments")
	pdf.Ln(12)
	pdf.SetFont("Arial", "", 12)
	for _, item := range users {
		count, err := a.Db.Comment.CountCommentsByUserId(item.Id)
		if err != nil {
			continue
		}
		line := item.DisplayName + " - comments: " + strconv.FormatUint(uint64(count), 10)
		pdf.Cell(0, 10, line)
		pdf.Ln(8)
	}

	// return pdf as file
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename=trevia_stats.pdf")
	err = pdf.Output(c.Writer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "PDF generation failed"})
		return
	}
}

func attractionRecordNumber(c *gin.Context) (int, error) {
	numberOfAttractionRecordsAsString := c.Query("attractions")
	number, err := strconv.Atoi(numberOfAttractionRecordsAsString)
	return number, err
}

func userRecordNumber(c *gin.Context) (int, error) {
	numberOfuserRecordsAsString := c.Query("users")
	number, err := strconv.Atoi(numberOfuserRecordsAsString)
	return number, err
}

func recordsNumber(c *gin.Context) (int, int) {
	defaultValue := 5
	numberOfAttractionRecords, attractionErr := attractionRecordNumber(c)
	numberOfUserRecords, userErr := userRecordNumber(c)
	if attractionErr != nil || userErr != nil {
		return defaultValue, defaultValue
	}
	return numberOfAttractionRecords, numberOfUserRecords
}
