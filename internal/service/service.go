package service

type ICreateMatrix interface {
	CreateMatrix(size int) error
}

type Service struct {
	ICreateMatrix
}

func NewService() *Service {
	return &Service{
		ICreateMatrix: NewMatrix(),
	}
}
