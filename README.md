Golang type's memory layout
============================================================

[![GoDoc](https://godoc.org/github.com/leeychee/mlayout?status.png)](https://godoc.org/github.com/leeychee/mlayout)
[![Build Status](https://travis-ci.org/leeychee/mlayout.svg?branch=master)](https://travis-ci.org/leeychee/mlayout)

Package mlayout try to describe golang type's memory layout, especially
struct's. According to language's implements, different order of fields use
different memory layout, so different memory use ratio sometimes.

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
```

So we could try to imporve the struct's use ratio of memory with this, by adjusting
struct's fields order.

## Installation

```go get -u github.com/leeychee/mlayout```

## Usage

For now, you can use the lib to get a type's use ratio of memory, and ascii image.

```go
type t1 struct {
	a bool
	b int
	c bool
}
l := mlayout.Layoutof(t1{})
fmt.Printf("t1's use ratio is %.2f", l.UseRatio())
fmt.Printf("t1's acsii image is: \n%s\n", l)
```

## TODO

Maybe, we could build a cmd tool to analysis structs, and give some advices.
