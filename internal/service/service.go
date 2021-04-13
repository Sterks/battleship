package service


type ICreateMatrix interface {
	CreateMatrix(rangeInt int) error
}

type Service struct {
	ICreateMatrix
}

func NewService() *Service {
	return &Service{}
}

