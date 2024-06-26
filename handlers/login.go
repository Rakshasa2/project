package handlers

import (
	"package/db"
	"net/http"
	"os"

	"time"

	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

//функция для создания этого токена
func createJWTToken(id uuid.UUID) (string, error) {
	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Minute * 20).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
//хендлер логина- отдает JWT токен в котором хранится ID пользователя в БД
func (h BaseHandler) LoginUser(c *gin.Context) {
	var loginRequest db.UserLoginData
	if err := c.BindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.db.GetUserByEmail(loginRequest.Email)
	if err != nil {
		log.Printf("Error retrieving user by email: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		log.Printf("Error comparing passwords: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}
	tokenString, err := createJWTToken(user.UserId)
	if err != nil {
		log.Printf("Error creating JWT token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"accessToken": tokenString})
}
