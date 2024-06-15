package v1

import (
	"os"
	"strconv"

	"smuggr.xyz/net-work/api/v1/routes"
	"smuggr.xyz/net-work/common/configurator"
	"smuggr.xyz/net-work/common/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var Config = &configurator.Config.API
var Logger = logger.NewCustomLogger("api/v1")

var DefaultRouter *gin.Engine
var DevicesInteractionsGroup *gin.RouterGroup

func SetupCors() {
	DefaultRouter.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			Logger.Debug(origin)
			return true
		},
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	}))

	DefaultRouter.SetTrustedProxies([]string{})
}

func Initialize() (*logger.MessageWrapper, chan *logger.MessageWrapper) {
	Config = &configurator.Config.API
	gin.SetMode(os.Getenv("GIN_MODE"))

	DefaultRouter = gin.Default()
	SetupCors()

	routes.Initialize(DefaultRouter)

	errCh := make(chan *logger.MessageWrapper)
	go func() {
		err := DefaultRouter.Run(":" + strconv.Itoa(int(Config.Port)))
		errCh <- logger.ErrUnexpected.Format(err.Error())
	}()

	Logger.Log(logger.MsgInitialized)

	return logger.MsgInitialized, errCh
}

func Cleanup() *logger.MessageWrapper {
	return nil
}
