Get Golang type's memory layout
============================================================

The lib will get golang type's memory layout. Different order of struct's fields
will use a different memory layout.

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

	fmt.Printf("t1:\n%s\n", memolayout.Layoutof(t1{}))
	fmt.Printf("t2:\n%s", memolayout.Layoutof(t2{}))
	// Output:
	// t1:
	// x_______
	// xxxxxxxx
	// x_______
	// t2:
	// xx______
	// xxxxxxxx
}
```
