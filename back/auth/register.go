package auth

import (
	"back/model"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type registerRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	DisplayName string `json:"displayName"`
}

type registerResponse struct {
	Token string `json:"token"`
	Id uint `json:"id"`
	IsModerator bool `json:"isModerator"`
	IsAdmin bool `json:"isAdmin"`
}

func (h *Handler) Register(c *gin.Context) {
	var request registerRequest

	c.BindJSON(&request)

	if request.Username == "" || request.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"Login or password is empty",
		})
		return
	}

	if !isUsernameValid(request.Username) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"Illegal login",
		})
		return
	}

	hashedPassword, err := h.HashPassword(request.Password); 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"Bad request",
		})
		return
	}
	
	var userRole model.Role
	userRole, err = h.Db.Role.SelectRoleByName("user")
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"message":"Service failure",
		})
		return
	}

	newUser := model.User {
		Username: request.Username,
		Password: hashedPassword,
		DisplayName: request.DisplayName,
		RoleId: userRole.Id,
	}

	// check if user exists in database
	_, err = h.Db.User.SelectByUsername(newUser.Username)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "This login already exists",
		})
		return
	}

	// insert user to database
	id, err := h.Db.User.CreateUser(newUser)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"message":"Service failure",
		})
		return
	}

	// response ok
	code, json, tokenString := h.login(request.Username, request.Password)
	if tokenString == nil {
		c.JSON(code, json)
		return
	}
	
	//awd

	response := registerResponse{
		Token: *tokenString,
		Id: id,
		IsModerator: json["role"] == "moderator",
		IsAdmin: json["role"] == "admin",
	}

	c.JSON(http.StatusOK, response)
}