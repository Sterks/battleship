package test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"
)

const (
	shippik = '\u25A1'
	hit     = "x"
)

func Test_location(t *testing.T) {
	g := make([][]string, 10)
	for i := 0; i < 10; i++ {
		g[i] = make([]string, 10)

		for j := range g[i] {
			if i == 0 && j == 0 {
				g[i][j] = string(shippik)
			} else if i == 0 && j == 1 {
				g[i][j] = string(shippik)
			} else if i == 1 && j == 0 {
				g[i][j] = string(shippik)
			} else if i == 1 && j == 1 {
				g[i][j] = string(shippik)
			} else {
				g[i][j] = " "
			}
		}
	}

	for m := range g {
		fmt.Println(g[m])
	}
}

func Test_coordinates(t *testing.T) {
	a := "5A"
	var mm Location
	mm = Location(a)
	col := mm.Column()
	row := mm.Row()
	println(row, col)
}

type Location string

func (l Location) Row() int {
	row := strings.Index("ABCDEFGHIJKLMNOPQRSTUVWXYZ", string(l[1:2]))
	return row
}

func (l Location) Column() int {
	column, _ := strconv.Atoi(string(l[:1]))
	return column - 1
}

type Board struct {
	size      int
	localtion [][]string
}

func Test_char(t *testing.T) {
	board := NewBoard()
	board.AddShip("1A 1B 1C, 9D 9E, 5F 6F 7F, 11F")

	for i := 0; i < len(board.localtion); i++ {
		fmt.Println(board.localtion[i])
	}

	cm := make(chan string)

	go func() {
		for {
			m, ok := <-cm
			if ok == false {
				break
			}
			board.Shot(m)
		}
	}()

	rangeCoord := "1A, 3E, 5G, 9C, 7A"
	cors := strings.Split(rangeCoord, ",")
	for _, value := range cors {
		value = strings.TrimSpace(value)
		cm <- value
		time.Sleep(10 * time.Second)
		for i := 0; i < len(board.localtion); i++ {
			fmt.Println(board.localtion[i])
		}
	}
	close(cm)

	//time.Sleep(1 *time.Minute)
}

func (b *Board) Shot(coord string) {
	a := Location(coord).Column()
	c := Location(coord).Row()

	for i := 0; i < len(b.localtion); i++ {
		for j := range b.localtion[i] {
			if a == i && c == j {
				if b.localtion[i][j] == string(shippik) {
					println(i, j)
					b.localtion[i][j] = hit
				} else {
					b.localtion[i][j] = "-"
				}
			}
		}
	}
}

func (b *Board) AddShip(coord string) {
	ships := strings.Split(coord, ",")
	for _, value := range ships {
		value := strings.TrimSpace(value)
		firstship := strings.Split(value, " ")
		for _, val := range firstship {
			sh := Location(val)
			for i := 0; i < 10; i++ {
				for j := range b.localtion[i] {
					if i == sh.Column() && j == sh.Row() {
						b.localtion[i][j] = string(shippik)
					}
				}
			}
		}
	}
}

func NewBoard() Board {
	var b Board
	rangeInt := 10
	g := make([][]string, rangeInt)
	for i := 0; i < 10; i++ {
		g[i] = make([]string, rangeInt)
	}
	b.localtion = g
	return b
}
