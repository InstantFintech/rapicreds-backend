package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"net/http"
	"rapicreds-backend/src/app/domain"
	"rapicreds-backend/src/app/util"
)

var googleOauth2Config = oauth2.Config{
	ClientID:     "590434445999-r3h4l926bnbi8b39e9u7n7jjkr3i4t41.apps.googleusercontent.com",
	ClientSecret: "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",
	RedirectURL:  "http://localhost:8080/auth/google/callback",
	Scopes:       []string{"openid", "profile", "email"},
	Endpoint:     google.Endpoint,
}

func GoogleLogin(c *gin.Context) {
	url := googleOauth2Config.AuthCodeURL("", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusFound, url)
}

func GoogleCallback(c *gin.Context) {
	code := c.DefaultQuery("code", "")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing code"})
		return
	}

	token, err := googleOauth2Config.Exchange(c, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to exchange token"})
		return
	}

	client := googleOauth2Config.Client(c, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get user info"})
		return
	}
	defer resp.Body.Close()

	var userInfo struct {
		ID    string `json:"id"`
		Email string `json:"email"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to parse user info"})
		return
	}

	// Buscar o crear usuario en la base de datos
	var user domain.User
	if err := db.Where("google_id = ?", userInfo.ID).First(&user).Error; err != nil {
		// Crear un nuevo usuario si no existe
		user = domain.User{
			Email:    userInfo.Email,
			GoogleID: userInfo.ID,
		}
		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}
	}

	// Generar un JWT token
	tokenStr, err := util.GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenStr})
}
