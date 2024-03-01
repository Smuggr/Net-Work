package routes

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"network/data/configuration"
	"network/data/errors"
	"network/data/messages"
	"network/data/models"
	"network/services/database"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)


func createToken(login string) (string, *errors.ErrorWrapper) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login": login,
		"exp":   time.Now().Add(time.Duration(configuration.Config.API.JWTLifespanMinutes) * time.Minute).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_TOKEN")))
	if err != nil {
		return "", errors.ErrSigningToken
	}

	return tokenString, nil
}

func AuthenticateUserHandler(c *gin.Context) {
    var user models.User
    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestPayload})
        return
    }

    existingUser := database.GetUser(database.DB, user.Login)
	if existingUser == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": errors.ErrUserNotFound})
        return
	}

    if err := database.AuthenticateUserPassword(existingUser, &user); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": errors.ErrInvalidCredentials})
        return
    }

    tokenString, err := createToken(user.Login)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func RegisterUserHandler(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestPayload})
		return
	}

	if err := database.RegisterUser(database.DB, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	tokenString, err := createToken(user.Login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": messages.MsgUserRegisterSuccess, "token": tokenString})
}

func UpdateUserHandler(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestPayload})
		return
	}

	if err := database.UpdateUser(database.DB, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": messages.MsgUserUpdateSuccess})
}

func RemoveUserHandler(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestPayload})
		return
	}

	if err := database.RemoveUser(database.DB, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": messages.MsgUserRemoveSuccess})
}


func GetAllUsersHandler(c *gin.Context) {
	users, err := database.GetLimitedUsers(database.DB, -1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrFetchingUsersFromDB})
		return
	}

    c.JSON(http.StatusOK, gin.H{
		"message": messages.MsgUsersFetchSuccess, 
		"users":   users,
	})
}

func GetLimitedUsersHandler(c *gin.Context) {
	limitStr := c.Query("limit")
    limit, err := strconv.Atoi(limitStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestPayload})
		return
    }

    users, err := database.GetLimitedUsers(database.DB, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrFetchingUsersFromDB})
		return
	}

    c.JSON(http.StatusOK, gin.H{
		"message": messages.MsgUsersFetchSuccess,
		"limit":   limit,
		"users":   users,
	})
}

func GetPaginatedUsersHandler(c *gin.Context) {
	pageStr := c.Query("page")
	pageSizeStr := c.Query("pageSize")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestPayload})
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestPayload})
		return
	}

	users, err := database.GetPaginatedUsers(database.DB, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrFetchingUsersFromDB})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"page":     page,
		"pageSize": pageSize,
		"users":    users,
	})
}
