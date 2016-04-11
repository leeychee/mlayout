package mlayout_test

import (
	"fmt"

	"github.com/leeychee/mlayout"
)

func ExampleLayoutof() {
	type t1 struct {
		a bool
		b int
		c bool
	}

	type t2 struct {
		a bool
		b bool
		c int
	}

	fmt.Printf("t1:\n%s\n", mlayout.Layoutof(t1{}))
	fmt.Printf("t2:\n%s", mlayout.Layoutof(t2{}))
	// Output:
	// t1:
	// x_______
	// xxxxxxxx
	// x_______
	// t2:
	// xx______
	// xxxxxxxx
}
