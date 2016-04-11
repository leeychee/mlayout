package mlayout

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
func TestLayoutBlankStruct(t *testing.T) {
	s := struct{}{}
	fmt.Printf("%s\n", Layoutof(s))
}
