package handlers

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"smuggr.xyz/net-work/common/logger"
	"smuggr.xyz/net-work/core/datastorer"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func createUserToken(login string) (string, *logger.MessageWrapper) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login": login,
		"exp":   time.Now().Add(time.Duration(APIConfig.JWTLifespanMinutes) * time.Minute).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_TOKEN")))
	if err != nil {
		return "", logger.ErrSigningToken
	}

	return tokenString, nil
}

func AuthenticateUserHandler(c *gin.Context) {
	var user datastorer.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": logger.ErrInvalidRequestPayload,
		})

		return
	}

	existingUser, _ := datastorer.GetUser(user.Login)
	if existingUser == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": logger.ErrResourceNotFound.Format(user.Login, logger.ResourceUser),
		})

		return
	}

	if err := datastorer.AuthenticateUserPassword(existingUser, user.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": logger.ErrInvalidCredentials,
		})

		return
	}

	tokenString, err := createUserToken(user.Login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": logger.MsgResourceAuthenticateSuccess.Format(user.Login, logger.ResourceUser),
		"token":   tokenString,
	})
}
func ValidateUserTokenHandler(c *gin.Context) {
	var requestBody struct {
		Login string `json:"login"`
		Token string `json:"token"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": logger.ErrInvalidRequestPayload,
		})

		return
	}

	login := requestBody.Login
	tokenString := requestBody.Token

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, logger.ErrSigningToken
		}

		return []byte(os.Getenv("SECRET_TOKEN")), nil
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": logger.ErrInvalidToken,
		})

		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["login"].(string) != login {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": logger.ErrInvalidToken,
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": logger.MsgResourceAuthenticateSuccess.Format(login, logger.ResourceUser),
		})

		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"error": logger.ErrInvalidToken,
	})
}

func RegisterUserHandler(c *gin.Context) {
	var user datastorer.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": logger.ErrInvalidRequestPayload,
		})

		return
	}

	if err := datastorer.RegisterUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	tokenString, err := createUserToken(user.Login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": logger.MsgResourceRegisterSuccess.Format(user.Login, logger.ResourceUser),
		"token":   tokenString,
	})
}

func UpdateUserHandler(c *gin.Context) {
	var user datastorer.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": logger.ErrInvalidRequestPayload,
		})

		return
	}

	if err := datastorer.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": logger.MsgResourceUpdateSuccess.Format(user.Login, logger.ResourceUser),
	})
}

func RemoveUserHandler(c *gin.Context) {
	login := c.Query("login")

	if login == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": logger.ErrInvalidRequestPayload,
		})

		return
	}

	user := datastorer.GetUser(login)
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": errors.ErrUserNotFound.Format(login)})
		return
	}

	if err := database.RemoveUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrRemovingUserFromDB.Format(login)})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": messages.MsgUserRemoveSuccess.Format(login)})
}

func GetUserHandler(c *gin.Context) {
	login := c.Param("login")

	user := database.GetUser(login)
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": errors.ErrUserNotFound.Format(login)})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": messages.MsgUserFetchSuccess.Format(login),
		"user":    user,
	})
}

func GetAllUsersHandler(c *gin.Context) {
	users, err := database.GetLimitedUsers(-1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrFetchingUsersFromDB})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": messages.MsgUsersFetchSuccess.Format(len(users)),
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

	users, err := database.GetLimitedUsers(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrFetchingUsersFromDB})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": messages.MsgUsersFetchSuccess.Format(len(users)),
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

	users, err := database.GetPaginatedUsers(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrFetchingUsersFromDB})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  messages.MsgUsersFetchSuccess.Format(len(users)),
		"page":     page,
		"pageSize": pageSize,
		"users":    users,
	})
}
