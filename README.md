Get Golang type's memory layout
============================================================

 Package memolayout try to describe a golang type's memo layout, especially struct's.
According to language's implements, different order of fields use different
memo layout, so different memo use ratio sometimes.
Here is a simple example, two structs with different fields order.

```go
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
```

## Usage

For now, you can use the lib to get a type's use ratio of memo, and ascii image.

```go
type t1 struct {
	a bool
	b int
	c bool
}
l := memolayout.Layoutof(t1{})
fmt.Printf("t1's use ratio is %.2f", l.UseRatio())
fmt.Printf("t1's acsii image is: \n%s\n", l)
```
