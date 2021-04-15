package service

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	shippik = '\u25A1'
	hit     = "x"
)

type Board struct {
	size     int
	Location [][]string
}

func NewBoard() *Board {
	return &Board{}
}

func (b *Board) CreateBoard(rangeInt int) error {
	b.Location = make([][]string, rangeInt)
	b.size = rangeInt
	for i := 0; i < rangeInt; i++ {
		b.Location[i] = make([]string, b.size)
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

func (b *Board) Shot(coord string) error {
	a := Location(coord).Column()
	c := Location(coord).Row()

	for i := 0; i < len(b.Location); i++ {
		for j := range b.Location[i] {
			if a == i && c == j {
				if b.Location[i][j] == string(shippik) {
					println(i, j)
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
