package jobcontroller

import "github.com/labstack/echo"

func (jc *jobController) FindHandler(e echo.Context) error {
	var (
		req = e.Request()
		ctx = req.Context()
	)

	return jc.jobUseCase.FindAllJobs(ctx).JSON(e)
}
