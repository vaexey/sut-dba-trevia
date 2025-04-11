package main

import (
	"back/config"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "my_secret_key"

func main() {
	
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group(config.Config.ApiPath)
	{
		api.GET("/hello", func(ctx *gin.Context){
			ctx.JSON(http.StatusOK, gin.H{"msg": "Hello world"})
		})
		
		api.POST("/login", handleLogin)

		api.Use(authMiddleware()) 
		{
			// TODO: auth endpoints 
		}
	}

	router.NoRoute(func(ctx *gin.Context){
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "not found"})
	})

	addr := ":" + strconv.Itoa(config.Config.Server.Port)
	router.Run(addr)
}

type loginRequestBody struct{
	Login string
	Password string
}

func handleLogin(c *gin.Context) {
	var loginRequestBody loginRequestBody
	
	if err := c.ShouldBindJSON(&loginRequestBody); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message" : "Something went wrong",
		})
		return
	}
   
	// TODO: retrieve user from the database, validate credentials, and determine role (change commented code block)

	 var role string
	// if loginRequestBody.Login == "admin" && loginRequestBody.Password == "password" {
	//  role = "admin"
	// } else if loginRequestBody.Login == "user" && loginRequestBody.Password == "password" {
	//  role = "user"
	// } else {
	//  c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	//  return
	// }
   
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": loginRequestBody.Login,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})
   
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
	 c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
	 return
	}
   
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		// parse token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrAbortHandler
			}
			return []byte(secretKey), nil
		})

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error" : "Unauthorized",
			})
			c.Abort();
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("claims", claims)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error" : "Unauthorized",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func adminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := c.MustGet("claims").(jwt.MapClaims)
		role := claims["role"].(string)

		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Forbidden",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
