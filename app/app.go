package app

import (
	"pandudpn/api/app/dbc"
	"pandudpn/api/app/dbc/sql"
	"pandudpn/api/src/controller"
	"pandudpn/api/src/controller/jobcontroller"
	"pandudpn/api/src/controller/visitorcontroller"
	"pandudpn/api/src/middleware"
	"pandudpn/api/src/repository"
	"pandudpn/api/src/repository/jobrepo"
	"pandudpn/api/src/repository/visitorrepo"
	"pandudpn/api/src/routes"
	"pandudpn/api/src/usecase"
	"pandudpn/api/src/usecase/jobuc"
	"pandudpn/api/src/usecase/visitoruc"
	"pandudpn/api/src/utils/config"
)

type app struct {
	dbc dbc.SqlDbc
	// repository
	jobRepo     repository.JobRepositoryInterface
	visitorRepo repository.VisitorRepositoryInterface
	// usecase
	jobUc     usecase.JobUseCaseInterface
	visitorUc usecase.VisitorUseCaseInterface
	// controller
	jobC     controller.JobControllerInterface
	visitorC controller.VisitorControllerInterface
}

func NewApp() *app {
	return &app{}
}

func (a *app) Register() {
	a.registerDatabase()
	a.registerRepo()
	a.registerUseCase()
	a.registerController()

	mdl := middleware.NewMiddleware(config.Logrus())

	router := routes.Route{
		Middleware:        mdl,
		JobController:     a.jobC,
		VisitorController: a.visitorC,
	}

	router.Router()
}

func (a *app) registerDatabase() {
	var d dbc.SqlDbc
	db := dbc.NewSqlConnection()

	d = &sql.Db{DB: db}
	a.dbc = d
}

// registerRepo used for inject all repository we have
func (a *app) registerRepo() {
	// job experiences
	job := jobrepo.NewJobRepository(a.dbc)
	a.jobRepo = job

	// visitor
	visitor := visitorrepo.NewVisitorRepository(a.dbc)
	a.visitorRepo = visitor
}

func (a *app) registerUseCase() {
	// job experiences
	job := jobuc.NewJobUseCase(a.jobRepo)
	a.jobUc = job

	// visitor
	visitor := visitoruc.NewVisitorUseCase(a.visitorRepo)
	a.visitorUc = visitor
}

func (a *app) registerController() {
	// Job experiences
	job := jobcontroller.NewJobController(a.jobUc)
	a.jobC = job

	// visitor
	visitor := visitorcontroller.NewVisitorHandler(a.visitorUc)
	a.visitorC = visitor
}
