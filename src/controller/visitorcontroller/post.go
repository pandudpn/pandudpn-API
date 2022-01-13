package visitorcontroller

import (
	"pandudpn/api/src/model"

	"github.com/labstack/echo"
)

func (v *VisitorHandler) NewVisitorHandler(c echo.Context) error {
	var (
		req     = c.Request()
		ctx     = req.Context()
		payload model.VisitorRequest
	)

	payload.Ip = c.RealIP()
	payload.UserAgent = req.UserAgent()

	return v.visitorUseCase.NewVisitor(ctx, &payload).JSON(c)
}
