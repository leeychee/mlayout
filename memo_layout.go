// Package memolayout try to describe a golang type's memo layout, especially struct's.
// According to language's implements, different order of fields use different
// memo layout, so different memo use ratio sometimes.
// Here is a simple example, two structs with different fields order.
//
// The layout is on Linux amd64.
//	type t1 struct {
//		b1 bool
//		s  string
//		b2 bool
//	}
//	// Layout of t1:
//	// x_______
//	// xxxxxxxx
//	// x_______
//
//	type t2 struct {
//		b1 bool
//		b2 bool
//		s  string
//	}
//	// Layout of t2:
//	// xx______
//	// xxxxxxxx
//
// As you can see, t1 use 24 byte, t2 use 16 byte, save more than 30% memo.
package memolayout

import (
	"bytes"
	"reflect"
)

type Layout [][]bool

// Layoutof will return a two-dimension array to describe the type's memo
// layout. It based on reflect, so it will get the real type of the interface.
func Layoutof(i interface{}) Layout {
	t := reflect.TypeOf(i)
	align := t.Align()
	size := int(t.Size())
	g := make([][]bool, size/align)
	for m := range g {
		line := make([]bool, align)
		for i := range line {
			line[i] = false
		}
		g[m] = line
	}
	layoutof(g, align, 0, t)
	return g
}

func layoutof(g [][]bool, align, offset int, t reflect.Type) {
	switch t.Kind() {
	case reflect.Struct:
		for m := 0; m < t.NumField(); m++ {
			f := t.Field(m)
			foffset := offset + int(f.Offset)
			layoutof(g, align, foffset, f.Type)
		}
	default:
		y := offset / align
		x := offset - y*align
		size := int(t.Size())
		for n := 0; n < size; n++ {
			if x >= align {
				x, y = 0, y+1
			}
			g[y][x] = true
			x++
		}
	}
}

// Use will return how much byte used in the memo.
func (g Layout) Use() (used int) {
	for _, l := range g {
		for _, p := range l {
			if p {
				used++
			}
		}
	}
	return
}

// UseRatio will return use ratio of the memo.
// Bigger is better.
func (g Layout) UseRatio() (ratio float32) {
	if len(g) < 1 {
		return 1.0
	}
	return float32(g.Use()) / float32(len(g)*len(g[0]))
}

// String will return a ascii image of the memo layout.
func (g Layout) String() string {
	if len(g) < 1 {
		return ""
	}
	var buf bytes.Buffer
	for i, line := range g {
		if i > 0 {
			buf.WriteByte('\n')
		}
		for _, point := range line {
			if point {
				buf.WriteByte('x')
			} else {
				buf.WriteByte('_')
			}
		}
	}
	return buf.String()
}
