package routes

import (
	"back/auth"
	"back/config"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type currentUserResponse struct {
	Token string 	`json:"token"`
	Id uint 		`json:"id"`
	RoleId uint 	`json:"roleId"`
}

func (a *Api) GetCurrentUser(c *gin.Context) {
	currentUserId := auth.ContextId(c);

	currentUser, err := a.Db.User.SelectById(*currentUserId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid request",
		})
		return
	}

	userRole, err := a.Db.Role.SelectById(currentUser.RoleId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid request",
		})
		return
	}

	tokenString, err := auth.CreateTokenString(currentUser.Username, currentUser.Id, userRole.Name, config.Config.SecretKey)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid request",
		})
		return
	}

	response := currentUserResponse{
		Token: tokenString,
		Id: currentUser.Id,
		RoleId: currentUser.RoleId,
	}

	c.JSON(http.StatusOK, response)
}