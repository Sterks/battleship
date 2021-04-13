package service

type Board struct{
	size int
	Location [][]int
}

func NewMatrix(location [][]int) *Board {
	return &Board{Location: location}
}

func (m *Board) CreateMatrix(rangeInt int) error {
	matrix := make([][]int, rangeInt)
	return nil
}

