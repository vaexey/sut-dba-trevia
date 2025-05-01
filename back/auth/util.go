package auth

import (
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (h *Handler) CompareHash(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func CtxIs(c *gin.Context) string {
	claims := c.MustGet("claims").(jwt.MapClaims)
	role := claims["role"].(string)
	
	return role
}

func CtxIsModerator(c *gin.Context) bool {
	role := CtxIs(c)
	return role == "moderator"
}

func CtxIsAdmin(c *gin.Context) bool {
	role := CtxIs(c)
	return role == "admin"
}

func CtxId(c *gin.Context) *uint {
	claims := c.MustGet("claims").(jwt.MapClaims)
	id := uint(claims["id"].(float64))
	return &id
}

func isUsernameValid(name string) bool {
	_, err := regexp.MatchString("^[a-zA-Z\\d_]+$", name)
	return err == nil
}