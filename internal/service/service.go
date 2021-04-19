package service

import (
	"battleship/internal/domain"
	"battleship/pkg/size"
)

type ICreateBoard interface {
	CreateBoard(size int) error
	ShowBoard() error
	AddShipInBoard(coordinates string) error
	Shot(coord string) error
	ShotResult(coord string) domain.Shot
	Clear() error
}

type IShot interface {


}

type Service struct {
	ICreateBoard ICreateBoard
	IShot IShot
}

type Deps struct {
	MapSize size.ISize
}

func NewService(deps Deps) *Service {
	return &Service{
		ICreateBoard: NewBoard(),
		IShot: NewShot(),
	}
}
