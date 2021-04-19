package service

import (
	"battleship/internal/domain"
	"fmt"
	"strconv"
	"strings"
)

const (
	shippik = '\u25A1'
	hit     = "x"
)

type Board struct {
	Size     int
	Location [][]string
}

func NewBoard() *Board {
	return &Board{}
}

func (b *Board) CreateBoard(rangeInt int) error {
	b.Location = make([][]string, rangeInt)
	b.Size = rangeInt
	for i := 0; i < rangeInt; i++ {
		b.Location[i] = make([]string, b.Size)
	}
	for i := 0; i < len(b.Location); i++ {
		for j := range b.Location[i] {
			b.Location[i][j] = " "
		}
	}

	return nil
}

func (b *Board) ShowBoard() error {
	for i := 0; i < len(b.Location); i++ {
		fmt.Println(b.Location[i])
	}
	return nil
}

type Location string

func (b *Board) AddShipInBoard(coordinates string) error {
	ships := strings.Split(coordinates, ",")
	for _, value := range ships {
		value := strings.TrimSpace(value)
		firstship := strings.Split(value, " ")
		for _, val := range firstship {
			sh := Location(val)
			for i := 0; i < 10; i++ {
				for j := range b.Location[i] {
					if i == sh.Column() && j == sh.Row() {
						b.Location[i][j] = string(shippik)
					}
				}
			}
		}
	}
	return nil
}

func (b *Board) ShotResult(coord string) domain.Shot {
	a := Location(coord).Column()
	c := Location(coord).Row()

	var sh domain.Shot
	for i := 0; i < len(b.Location); i++ {
		for j := range b.Location[i] {

			// 6E = 6
			if a == i && c == j {
				// 5E
				isLive := b.Check(i, j)
				sh.Destroy = isLive

				if b.Location[i][j] == hit {
					sh.Knock = true
				}
			}
		}
	}
	fmt.Println(sh)
	return sh
}

func (b *Board) Check(i int, j int) bool {
	var isLive bool
	if i > 0 {
		if b.Location[i-1][j] == string(shippik) {
			//result["destroy"] = false
			//s.Board.Location[i-1][j] = "O"
		} else {
			//s.Board.Location[i-1][j] = "M"
			isLive = true
		}
	}
	//2  6D
	if j > 0 {
		if b.Location[i][j-1] == string(shippik) {
			//result["destroy"] = false
			//s.Board.Location[i][j-1] = "O"
		} else {
			//s.Board.Location[i][j-1] = "M"
			isLive = true
		}
	}
	//3 7E
	if i < len(b.Location) - 1 {
		if b.Location[i+1][j] == string(shippik) {
			//result["destroy"] = false
			//s.Board.Location[i+1][j] = "O"
		} else {
			//s.Board.Location[i+1][j] = "M"
			isLive = true
		}
	}
	//4 6F
	if j < len(b.Location) - 1 {
		if b.Location[i][j+1] == string(shippik) {
			//s.Board.Location[i][j+1] = "O"
		} else {
			//s.Board.Location[i][j+1] = "M"
			isLive = true
		}
	}
	//}
	return isLive
}

func (b *Board) Shot(coord string) error {
	a := Location(coord).Column()
	c := Location(coord).Row()

	for i := 0; i < len(b.Location); i++ {
		for j := range b.Location[i] {
			if a == i && c == j {
				if b.Location[i][j] == string(shippik) {
					fmt.Printf("Координаты выстрела %v: вертикаольно %v\n", coord, i )
					fmt.Printf("					 горизонтально %v\n", j)
					b.Location[i][j] = hit
				} else {
					b.Location[i][j] = "-"
				}
			}
		}
	}
	return nil
}

func (b *Board) Clear() error {
	for i := 0; i < len(b.Location); i++ {
		for j := range b.Location[i] {
			b.Location[i][j] = ""
		}
	}
	return nil
}

func (l Location) Row() int {
	row := strings.Index("ABCDEFGHIJKLMNOPQRSTUVWXYZ", string(l[1:2]))
	return row
}

func (l Location) Column() int {
	column, _ := strconv.Atoi(string(l[:1]))
	return column - 1
}
