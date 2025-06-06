package auth

import (
    "regexp"
    "time"

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

func ContextRole(c *gin.Context) string {
    claims := c.MustGet("claims").(jwt.MapClaims)
    role := claims["role"].(string)
    
    return role
}

func ContextIsModerator(c *gin.Context) bool {
    role := ContextRole(c)
    return role == "moderator" || role == "admin"
}

func ContextIsAdmin(c *gin.Context) bool {
    role := ContextRole(c)
    return role == "admin"
}

func ContextId(c *gin.Context) *uint {
    claims := c.MustGet("claims").(jwt.MapClaims)
    id := uint(claims["id"].(float64))
    return &id
}

func CreateTokenString(username string, userId uint, userRole string, secretKey string) (string, error) {
    // create token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username" : username,
        "id": userId,
        "role": userRole,
        "exp": time.Now().Add(time.Hour * 1).Unix(),    
    })

    tokenstring, err := token.SignedString([]byte(secretKey))
    return tokenstring, err
}

func isUsernameValid(name string) bool {
    _, err := regexp.MatchString("^[a-zA-Z\\d_]+$", name)
    return err == nil
}
