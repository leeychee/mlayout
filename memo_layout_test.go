package memolayout

import (
	"fmt"
	"testing"
)

func TestLayoutofInt(t *testing.T) {
	var i int
	fmt.Printf("%s\n", Layoutof(i))
}
func TestLayoutofBool(t *testing.T) {
	var b bool
	fmt.Printf("%s\n", Layoutof(b))
}
func TestLayoutofFloat32(t *testing.T) {
	var f float32
	fmt.Printf("%s\n", Layoutof(f))
}
func TestLayoutofFloat64(t *testing.T) {
	var f float64
	fmt.Printf("%s\n", Layoutof(f))
}
func TestLayoutofSlice(t *testing.T) {
	var s []int
	fmt.Printf("%s\n", Layoutof(s))
}
func TestLayoutofMapStringString(t *testing.T) {
	var s map[string]string
	fmt.Printf("%s\n", Layoutof(s))
}
func TestLayoutofMapStringInt(t *testing.T) {
	var s map[string]int
	fmt.Printf("%s\n", Layoutof(s))
}
func TestLayoutStruct1(t *testing.T) {
	s := struct{}{}
	fmt.Printf("%s\n", Layoutof(s))
}
func TestLayoutStruct2(t *testing.T) {
	s := struct{ int }{1}
	fmt.Printf("%s\n", Layoutof(s))
}
func TestLayoutStruct3(t *testing.T) {
	s := struct {
		b bool
		i int
	}{
		true,
		1,
	}
	fmt.Printf("%s\n", Layoutof(s))
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
	fmt.Printf("%s\n", Layoutof(s))
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
	fmt.Printf("%s\n", Layoutof(s))
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
	fmt.Printf("%s\n", Layoutof(s))
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
	fmt.Printf("%s\n", Layoutof(s))
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
	fmt.Printf("%s\n", Layoutof(s))
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
	fmt.Printf("%s\n", Layoutof(s))
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
	fmt.Printf("%s\n", Layoutof(s))
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
	fmt.Printf("%s\n", Layoutof(s))
}
