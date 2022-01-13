package repository

import "pandudpn/api/src/model"

type JobRepositoryInterface interface {
	FindAllJobs() ([]*model.Job, error)
}

type VisitorRepositoryInterface interface {
	FindVisitorByIpAndUserAgent(ip, userAgent string) (*model.Visitor, error)
	NewVisitor(visitor *model.Visitor) error
	UpdateVisitor(visitor *model.Visitor) error
}
