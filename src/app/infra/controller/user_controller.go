package controller

import (
	"fmt"
	"net/http"
	"rapicreds-backend/src/app/domain"
	"rapicreds-backend/src/app/util"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	cookie, err := c.Cookie("session_token")
	if err != nil {
		c.JSON(http.StatusNotFound, domain.Response{Message: "Not found cookie"})
		return
	}

	tokenString := cookie
	claims := &jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return util.JwtKey, nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusNotFound, domain.Response{Message: "Not valid token"})
		return
	}

	userID := claims.Subject

	var user domain.User
	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found user"})
		return
	}

	c.JSON(http.StatusOK, user)
	return
}

func IsVerified(c *gin.Context) {
	cookie, err := c.Cookie("session_token")
	if err != nil {
		c.JSON(http.StatusNotFound, domain.Response{Message: "Not found cookie"})
		return
	}

	tokenString := cookie
	claims := &jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return util.JwtKey, nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusNotFound, domain.Response{Message: "Not valid token"})
		return
	}

	userID := claims.Subject

	var user domain.User
	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found user"})
		return
	}

	isVerified := "no"
	if user.IsVerified() {
		isVerified = "yes"
	}
	c.JSON(http.StatusOK, domain.Response{Message: isVerified})
	return
}

func UpdateUser(c *gin.Context) {
	cookie, err := c.Cookie("session_token")
	if err != nil {
		c.JSON(http.StatusNotFound, domain.Response{Message: "Not found cookie"})
		return
	}

	var bodyUser domain.User
	if err := c.ShouldBindJSON(&bodyUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user"})
		return
	}

	tokenString := cookie
	claims := &jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return util.JwtKey, nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusNotFound, domain.Response{Message: "Not valid token"})
		return
	}

	userID := claims.Subject

	var user domain.User
	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found user"})
		return
	}

	user.Name = bodyUser.Name
	user.LastName = bodyUser.LastName
	user.Sex = bodyUser.Sex
	user.Document = bodyUser.Document

	db.Updates(user)

	c.JSON(http.StatusOK, domain.Response{Message: "User updated"})
	return
}
