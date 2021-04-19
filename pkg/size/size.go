package size

type ISize interface {
	GetSize() int
	SetSize(size int)
}

type Size struct {
	Size int
}

func NewSize(size int) *Size {
	return &Size{
		Size: size,
	}
}
func(s *Size) SetSize(size int) {
	s.SetSize(size)
}

func (s *Size) GetSize() int {
	return s.Size
}