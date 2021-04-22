package gotest

import (
	"fmt"
	"github.com/ainsleyclark/go-test/interfaces"
	mocks "github.com/ainsleyclark/go-test/mocks/interfaces"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ShapeTestSuite struct {
	suite.Suite
}

// Assert testing has begun.
func TestShape(t *testing.T) {
	suite.Run(t, new(ShapeTestSuite))
}

// Mocks
//
// Using mocks allows use to change the way interfaces work through our testing.
// We can use stubs to change the output of our interfaces in separate
// packages.

// A SUCCESS mock implementation of our shape.
type mockGetSize struct{}

func (r mockGetSize) Area() (float64, error) {
	return 20.0, nil
}

// A ERROR mock implementation of our shape.
type mockGetSizeErr struct{}

func (r mockGetSizeErr) Area() (float64, error) {
	return 0, fmt.Errorf("shape error")
}

func (t *ShapeTestSuite) TestMyStruct_GetSize() {
	tt := map[string]struct {
		input interfaces.Shape
		want  interface{}
	}{
		"Success": {
			input: mockGetSize{},
			want:  20.0,
		},
		"Error": {
			input: mockGetSizeErr{},
			want:  "shape error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			got, err := test.input.Area()
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}
			t.Equal(test.want, got)
		})
	}
}

// Using Mockery
//
// Using mockery we can automatically generate our mocks and use a function
// to output the expected args.
func (t *ShapeTestSuite) TestMyStruct_GetSize_Mock() {
	tt := map[string]struct {
		input func(shape *mocks.Shape)
		want  interface{}
	}{
		"Success": {
			input: func(shape *mocks.Shape) {
				shape.On("Area").Return(20.0, nil)
			},
			want:  20.0,
		},
		"Error": {
			input: func(shape *mocks.Shape) {
				shape.On("Area").Return(0.0, fmt.Errorf("shape error"))
			},
			want:  "shape error",
		},
	}

	for name, test := range tt {
		t.Run(name, func() {
			s := &mocks.Shape{}
			test.input(s)
			got, err := s.Area()
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}
			t.Equal(test.want, got)
		})
	}
}
