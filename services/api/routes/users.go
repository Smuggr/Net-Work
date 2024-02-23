package routes

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"network/data/errors"
	"network/data/messages"
	"network/data/models"
	"network/services/database"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)


func createToken(login string) (string, *errors.ErrorWrapper) {
	jwt_token_lifespan_minutes, err := strconv.Atoi(os.Getenv("API_JWT_TOKEN_LIFESPAN"))
	if err != nil {
		return "", errors.ErrCreatingToken
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login": login,
		"exp":   time.Now().Add(time.Duration(jwt_token_lifespan_minutes) * time.Minute).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_TOKEN")))
	if err != nil {
		return "", errors.ErrSigningToken
	}

	return tokenString, nil
}

func AuthenticateUser(c *gin.Context) {
    var user models.User
    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestPayload.Key})
        return
    }

    var existingUser models.User
    if database.DB.Where("login = ?", user.Login).First(&existingUser).Error != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": errors.ErrUserNotFound.Key})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": errors.ErrInvalidCredentials.Key})
        return
    }

    tokenString, err := createToken(user.Login)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Key})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestPayload.Key})
		return
	}

	if err := database.RegisterUser(database.DB, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Key})
		return
	}

	tokenString, err := createToken(user.Login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Key})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": messages.MsgUserRegisterSuccess.Key, "token": tokenString})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestPayload.Key})
		return
	}

	if err := database.UpdateUser(database.DB, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Key})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": messages.MsgUserUpdateSuccess.Key})
}