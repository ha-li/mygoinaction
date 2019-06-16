package main

// shows
// 1. how polymorphism can be implemented in go using interfaces
// 2. Type Asserts using construct: interface{}
import (
	"fmt"
	"math"
)

// 3 different types that are Shapes
// 1 is Circle
type Circle struct {
	Radius float64
}

// a functino of Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// 2 is Triangle
type Triangle struct {
	A, B, C float64
}

// functions of Triangle are Area() and Angle
func (t Triangle) Area() float64 {
	p := (t.A + t.B + t.C)/2
	return math.Sqrt(p * (p-t.A) * (p-t.B) * (p - t.C))
}

func (t Triangle) Angle() []float64 {
	return []float64{ 12.0, 90, 78}
}


// 3 is Square
type Square struct {
	A float64
}

// the area of square
func (s Square) Area() float64 {
	return s.A * s.A
}

// an interface, note that all 3 types implement this interface
type Shape interface {
	Area() float64
}


// a function that does a Type Assert and only calls Angle on triangle types
func PrintAngles(y interface{}) {

	// Type Assert construct is v.(T)
	// where v is the instance, T is the tested interface
	// and will return 2 values:
	// 1, the instance as that tested type, if successful, otherwise nil
	// 2, a boolean if the conversion was successful
	if t, ok := y.(Triangle); ok  {
		fmt.Println("angles ",  t.Angle())
	} else {
		fmt.Println ( "not a triangle ", ok)
	}
}

func main() {
	// an area of Shape types
	shapes := []Shape {
		Circle {1.0},
		Square{3.0},
		Triangle {1, 2.0, 1.4},
	}

	for _, r := range shapes {
		fmt.Println(r.Area())
		PrintAngles(r)
	}
}