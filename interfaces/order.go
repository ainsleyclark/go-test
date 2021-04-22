package interfaces

// Shape An interface is a collection of method signatures that a Type can implement
// (using methods). Hence interface defines (not declares) the behavior of the
// object (of the type Type).
//
// For example, a Dog can walk and bark. If an interface defines method signatures for
// walk and bark while Dog implements walk and bark methods, then Dog is said to
// implement that interface.
//
// The primary job of an interface is to provide only method signatures consisting of
// the method name, input arguments and return types. It is up to a Type
// (e.g. struct type) to declare methods and implement them.

// Notes:
// 	- Powerful interfaces are single method, many types can implement them.
//  - We can say the Rect implements the Shape interface, we do not need
//    to use the implements key word as it is implicitly implemented.

// Shape is an interface the declares methods for structs to implement.
type Shape interface {
	Area() (float64, error)
}

// Rect - A struct or type that implements the interface.
type Rect struct {
	Width  float64
	Height float64
}

func (r Rect) Area() (float64, error) {
	return r.Width * r.Height, nil
}
