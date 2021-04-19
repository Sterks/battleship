package service

import (
	"battleship/internal/domain"
	"fmt"
)

type Shot struct {
	Coord string
	Board *Board
}

func NewShot() *Shot {
	return &Shot{}
}

func (s *Shot) Shot(coord string) error {
	a := Location(coord).Column()
	c := Location(coord).Row()

	i2 := len(s.Board.Location)
	for i := 0; i < i2; i++ {
		for j := range s.Board.Location[i] {
			if a == i && c == j {
				if s.Board.Location[i][j] == string(shippik) {
					fmt.Printf("Координаты выстрела %v: вертикаольно %v\n", coord, i )
					fmt.Printf("					 горизонтально %v\n", j)
					s.Board.Location[i][j] = hit
				} else {
					s.Board.Location[i][j] = "-"
				}
			}
		}
	}
	return nil
}

func (s *Shot) ShotResult() domain.Shot {
	a := Location(s.Coord).Column()
	c := Location(s.Coord).Row()

	var sh domain.Shot
	for i := 0; i < len(s.Board.Location); i++ {
		for j := range s.Board.Location[i] {

			// 6E = 6
			if a == i && c == j {
				// 5E
				isLive := s.Check(i, j)
				sh.Destroy = isLive

				if s.Board.Location[i][j] == hit {
					sh.Knock = true
				}
			}
		}
	}
	fmt.Println(sh)
	return sh
}

func (s *Shot) Check(i int, j int) bool {
	var isLive bool
		if i > 0 {
			if s.Board.Location[i-1][j] == string(shippik) {
				//result["destroy"] = false
				//s.Board.Location[i-1][j] = "O"
			} else {
				//s.Board.Location[i-1][j] = "M"
				isLive = true
			}
		}
		//2  6D
		if j > 0 {
			if s.Board.Location[i][j-1] == string(shippik) {
				//result["destroy"] = false
				//s.Board.Location[i][j-1] = "O"
			} else {
				//s.Board.Location[i][j-1] = "M"
				isLive = true
			}
		}
		//3 7E
		if i < len(s.Board.Location) - 1 {
			if s.Board.Location[i+1][j] == string(shippik) {
				//result["destroy"] = false
				//s.Board.Location[i+1][j] = "O"
			} else {
				//s.Board.Location[i+1][j] = "M"
				isLive = true
			}
		}
		//4 6F
		if j < len(s.Board.Location) - 1 {
			if s.Board.Location[i][j+1] == string(shippik) {
				//s.Board.Location[i][j+1] = "O"
			} else {
				//s.Board.Location[i][j+1] = "M"
				isLive = true
			}
		}
	//}
	return isLive
}