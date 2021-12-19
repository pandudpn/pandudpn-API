package app

import (
	"pandudpn/api/app/dbc"
	"pandudpn/api/app/dbc/sql"
	"pandudpn/api/src/repository"
	"pandudpn/api/src/repository/jobrepo"
	"pandudpn/api/src/usecase"
	"pandudpn/api/src/usecase/jobuc"
)

type app struct {
	dbc dbc.SqlDbc
	// repository
	jobRepo repository.JobRepositoryInterface
	// usecase
	jobUc usecase.JobUseCaseInterface
}

func NewApp() *app {
	return &app{}
}

func (a *app) Register() {
	a.registerDatabase()
	a.registerRepo()
	a.registerUseCase()
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
}

func (a *app) registerUseCase() {
	// job experiences
	job := jobuc.NewJobUseCase(a.jobRepo)
	a.jobUc = job
}
