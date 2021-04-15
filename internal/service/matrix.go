package service

type Board struct {
	size     int
	Location [][]int
}

func NewMatrix() *Board {
	return &Board{}
}

func (m *Board) CreateMatrix(rangeInt int) error {
	m.Location = make([][]int, rangeInt)
	for i := 0; i < rangeInt; i++ {
		m.Location[i] = make([]int, m.size)

		for j := range m.Location[i] {
			m.Location[i][j] = 5
		}
	}
	return nil
}
