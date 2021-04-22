package gotest

import (
	"github.com/ainsleyclark/go-test/interfaces"
)

type MyStruct struct {
	Shape interfaces.Shape
}

func (m *MyStruct) GetArea() (float64, error) {
	a, err := m.Shape.Area()
	if err != nil {
		return 0, err
	}
	return a, nil
}
