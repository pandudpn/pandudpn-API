package visitoruc

import (
	"net/http"

	"pandudpn/api/src/repository"
)

type visitorRes int

const (
	successInsert visitorRes = iota
	successUpdate
	errInsert
	errQuery
	errUpdate
)

type VisitorUseCase struct {
	visitorRepo repository.VisitorRepositoryInterface
}

func NewVisitorUseCase(vr repository.VisitorRepositoryInterface) *VisitorUseCase {
	return &VisitorUseCase{
		visitorRepo: vr,
	}
}

func (v visitorRes) StatusCode() int {
	return [...]int{
		http.StatusCreated,
		http.StatusOK,
		http.StatusInternalServerError,
		http.StatusInternalServerError,
		http.StatusInternalServerError,
	}[v]
}

func (v visitorRes) Message() string {
	errGlobalMessage := "Something went wrong, please try again later"

	return [...]string{
		"New IP success to stamp",
		"IP visit updated",
		errGlobalMessage,
		errGlobalMessage,
		errGlobalMessage,
	}[v]
}
