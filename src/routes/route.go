package routes

import (
	"fmt"

	"pandudpn/api/src/controller"
	"pandudpn/api/src/middleware"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

type Route struct {
	Middleware        middleware.MiddlewareInterface
	JobController     controller.JobControllerInterface
	VisitorController controller.VisitorControllerInterface
}

func NewRouter(jc controller.JobControllerInterface, vc controller.VisitorControllerInterface) *Route {
	return &Route{
		JobController:     jc,
		VisitorController: vc,
	}
}

func (r *Route) Router() {
	e := echo.New()
	e.HideBanner = true

	e.Use(echo.WrapMiddleware(r.Middleware.Logger))

	e.GET("/jobs", r.JobController.FindHandler)

	e.POST("/visitors", r.VisitorController.NewVisitorHandler)

	port := fmt.Sprintf(":%d", viper.GetInt("PORT"))
	e.Start(port)
}
