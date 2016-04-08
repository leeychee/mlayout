Get Golang type's memory layout
============================================================

The lib will get golang type's memory layout. The order of struct's fields will
use a different memory layout.

## Usage

```go
func ExampleLayout() {
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

	fmt.Printf("%s\n\n", memolayout.Layoutof(t1{}))
	fmt.Printf("%s", memolayout.Layoutof(t2{}))
	// Output:
	// x_______
	// xxxxxxxx
	// x_______
	//
	// xx______
	// xxxxxxxx
}
```
