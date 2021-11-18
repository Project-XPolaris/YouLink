package httpapi

import (
	"github.com/allentom/haruka"
	"github.com/allentom/haruka/middleware"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

var Logger = log.New().WithFields(log.Fields{
	"scope": "Application",
})

func GetEngine() *haruka.Engine {
	e := haruka.NewEngine()
	e.UseCors(cors.AllowAll())
	e.UseMiddleware(middleware.NewLoggerMiddleware())
	e.UseMiddleware(middleware.NewPaginationMiddleware("page", "pageSize", 1, 20))
	e.Router.POST("/program/template", createProgramHandler)
	e.Router.POST("/program/instance", newInstanceProgram)
	e.Router.POST("/process/run", runProcessHandler)
	e.Router.GET("/functions", functionListHandler)
	e.Router.POST("/callback", outputCallbackHandler)
	e.Router.POST("/register", registerFunctionHandler)
	e.Router.AddHandler("/notification", notificationSocketHandler)
	e.UseMiddleware(&AuthMiddleware{})
	return e
}
