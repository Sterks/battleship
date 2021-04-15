package service

type ICreateBoard interface {
	CreateBoard(size int) error
	ShowBoard() error
	AddShipInBoard(coordinates string) error
	Shot(coord string) error
	Clear() error
}

type Service struct {
	ICreateBoard
}

func NewService() *Service {
	return &Service{
		ICreateBoard: NewBoard(),
	}
}
