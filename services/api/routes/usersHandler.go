package routes

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"network/services/database"
	"network/utils/configuration"
	"network/utils/errors"
	"network/utils/messages"
	"network/utils/models"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func createUserToken(login string) (string, *errors.ErrorWrapper) {
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
		log.Debug(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestPayload})
		return
	}

	existingUser := database.GetUser(user.Login)
	if existingUser == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": errors.ErrUserNotFound.Format(user.Login)})
		return
	}

	if err := database.AuthenticateUserPassword(existingUser, user.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": errors.ErrInvalidCredentials})
		return
	}

	tokenString, err := createUserToken(user.Login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": messages.MsgUserAuthenticateSuccess.Format(user.Login),
		"token":   tokenString,
	})
}
func ValidateUserTokenHandler(c *gin.Context) {
	var requestBody struct {
		Login string `json:"login"`
		Token string `json:"token"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		log.Debug("error parsing request body: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	login := requestBody.Login
	tokenString := requestBody.Token

	log.Debug("parsing token")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Debug("invalid signing method")
			return nil, errors.ErrSigningToken
		}

		return []byte(os.Getenv("SECRET_TOKEN")), nil
	})

	if err != nil {
		log.Debug("error parsing token: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidToken})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["login"].(string) != login {
			log.Debug("token does not match login")
			c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidToken})
			return
		}

		log.Debug("token validated successfully")
		c.JSON(http.StatusOK, gin.H{"message": messages.MsgUserAuthenticateSuccess.Format(login)})
		return
	}

	log.Debug("invalid token")
	c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidToken})
}

func RegisterUserHandler(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestPayload})
		return
	}

	if err := database.RegisterUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	tokenString, err := createUserToken(user.Login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": messages.MsgUserRegisterSuccess.Format(user.Login), "token": tokenString})
}

func UpdateUserHandler(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestPayload})
		return
	}

	if err := database.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": messages.MsgUserUpdateSuccess.Format(user.Login)})
}

func RemoveUserHandler(c *gin.Context) {
	login := c.Query("login")

	if login == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestPayload})
		return
	}

	user := database.GetUser(login)
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
