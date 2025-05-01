package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token string `json:"token"`
	Id uint `json:"id"`
	IsModerator bool `json:"isModerator"`
	IsAdmin bool `json:"isAdmin"`
}


func (h *Handler) login(username string, password string) (int, gin.H, *string) {
	if username == "" || password == "" {
		return http.StatusUnauthorized, gin.H{ "message": "Invalid credentials1" }, nil
	}

	dbUser, err := h.Db.User.SelectByUsername(username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusUnauthorized, gin.H{ "message":"Invalid credentials2" }, nil
	}

	if err != nil {
		return http.StatusServiceUnavailable, gin.H{ "message":"Service failure" }, nil
	}

	if !h.CompareHash(password, dbUser.Password) || dbUser.Username != username {
		return http.StatusUnauthorized, gin.H{ "message":"Invalid credentials3" }, nil
	}

	var role string = "user"
	moderatorRole, moderatorErr := h.Db.Role.SelectRoleByName("moderator")
	adminRole, adminErr := h.Db.Role.SelectRoleByName("admin")
	if moderatorErr != nil || adminErr != nil {
		return http.StatusServiceUnavailable, gin.H{ "message":"Service failure", }, nil
	}

	if dbUser.RoleId == moderatorRole.Id{
		role = "moderator"
	} else if (dbUser.RoleId == adminRole.Id){
		role = "admin"
	}

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username" : username,
		"id": dbUser.Id,
		"role": role,
		"exp": time.Now().Add(time.Hour * 1).Unix(),	
	})

	// sign and get encoded token as string
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return http.StatusServiceUnavailable, gin.H { "message" : "Service failure" }, nil
	}

	return http.StatusOK, gin.H{ "message" : "success" , "id": dbUser.Id, "role": role}, &tokenString
}

func (h *Handler) Login(c *gin.Context) {
	var request loginRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "message": err.Error() })
		return
	}

	code, json, tokenString := h.login(request.Username, request.Password)
	if tokenString == nil {
		c.JSON(code, json)
		return
	}

	isModerator := false
	isAdmin := false

	if	json["role"] == "moderator" || json["role"] == "admin" {
		isModerator = true
	}
	if (json["role"] == "admin") {
		isAdmin = true
	}

	response := loginResponse{
		Token: *tokenString,
		Id: json["id"].(uint),
		IsModerator: isModerator,
		IsAdmin: isAdmin,
	}

	c.JSON(http.StatusOK, response)
}