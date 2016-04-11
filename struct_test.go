package memolayout

import (
	"fmt"
	"testing"
)

func TestLayoutStruct2(t *testing.T) {
	s := struct{ int }{1}
	printStructMemoLayout(s)
}
func TestLayoutStruct3(t *testing.T) {
	s := struct {
		b bool
		i int
	}{
		true,
		1,
	}
	printStructMemoLayout(s)
}
func TestLayoutStruct4(t *testing.T) {
	s := struct {
		b1 bool
		b2 bool
		i  int
	}{
		true,
		false,
		1,
	}
	printStructMemoLayout(s)
}
func TestLayoutStruct5(t *testing.T) {
	s := struct {
		a bool
		b int16
		c []int
	}{
		true,
		int16(1),
		[]int{1},
	}
	printStructMemoLayout(s)
}
func TestLayoutStruct6(t *testing.T) {
	s := struct {
		a bool
		b int16
		c []int
		d struct{ int }
	}{
		true,
		int16(1),
		[]int{1},
		struct{ int }{1},
	}
	printStructMemoLayout(s)
}
func TestLayoutStruct7(t *testing.T) {
	type st struct {
		a bool
		b int16
		c []int
		d *st
	}
	s := st{
		true,
		int16(1),
		[]int{1},
		&st{},
	}
	printStructMemoLayout(s)
}
func TestLayoutStruct71(t *testing.T) {
	type st struct {
		b int16
		a bool
		c []int
		d *st
	}
	s := st{
		int16(1),
		true,
		[]int{1},
		&st{},
	}
	printStructMemoLayout(s)
}
func TestLayoutStruct8(t *testing.T) {
	type t1 struct {
		a bool
		b int16
		c []int
	}
	type t2 struct {
		t1
		a string
	}
	s := t2{
		t1{
			true,
			int16(0),
			[]int{},
		},
		"name",
	}
	printStructMemoLayout(s)
}
func TestLayoutStruct9(t *testing.T) {
	type t1 struct {
		a bool
		b int16
		c []int
	}
	type t2 struct {
		a string
		t1
	}
	s := t2{
		"name",
		t1{
			true,
			int16(0),
			[]int{},
		},
	}
	printStructMemoLayout(s)
}
func TestLayoutStruct10(t *testing.T) {
	type st struct {
		b int16
		a bool
		c byte
		d bool
		e string
		f bool
	}
	s := st{
		int16(1),
		true,
		'\n',
		false,
		"n",
		true,
	}
	printStructMemoLayout(s)
}

func printStructMemoLayout(i interface{}) {
	l := Layoutof(i)
	fmt.Printf("%s\n", l)
	fmt.Printf("Use Ration: %.2f\n", l.UseRatio())
}
