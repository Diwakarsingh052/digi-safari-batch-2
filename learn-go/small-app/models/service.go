package models

//go:generate mockgen -source service.go -destination mockmodels/service_mock.go -package mockmodels
type Service interface {
	CreateUser(n NewUser) (User, error)
}

type Store struct {
	Service
}

func NewStore(s Service) Store {
	return Store{Service: s}
}
