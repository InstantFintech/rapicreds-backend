package controller

import (
	"fmt"
	"net/http"
	"rapicreds-backend/src/app/domain"
	"rapicreds-backend/src/app/domain/constants"
	"rapicreds-backend/src/app/util"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateLoan(c *gin.Context) {
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

	var userLoan domain.UserLoan
	if err := c.ShouldBindJSON(&userLoan); err != nil {
		fmt.Printf("Error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid userLoan"})
		return
	}

	fmt.Printf("UserLoan: %v", userLoan)
	user.Loans = append(user.Loans, userLoan)
	db.Updates(user)

	c.JSON(http.StatusOK, user)
	return
}

func GetLoan(c *gin.Context) {
	status := constants.LoanStatusPending
	statusQueryParam := constants.LoanStatus(c.Query("status"))

	if statusQueryParam == constants.LoanStatusApproved {
		status = constants.LoanStatusApproved
	}

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

	db.Table("user_loans").
		Where("user_id = ? AND status = ?", userID, status).
		Find(&user.Loans)

	c.JSON(http.StatusOK, user.Loans)
	return
}
