package memolayout

import (
	"bytes"
	"reflect"
)

type Layout [][]byte

func Layoutof(i interface{}) Layout {
	t := reflect.TypeOf(i)
	align := t.Align()
	size := int(t.Size())
	//fmt.Printf("T: %-10sS: %d\tA: %d\n", t.Kind(), size, align)
	g := make([][]byte, size/align)
	for m := range g {
		line := make([]byte, align)
		for i := range line {
			line[i] = '_'
		}
		g[m] = line
	}
	layoutof(g, align, 0, t)
	return g
}

func layoutof(g [][]byte, align, offset int, t reflect.Type) {
	switch t.Kind() {
	case reflect.Struct:
		for m := 0; m < t.NumField(); m++ {
			f := t.Field(m)
			//falign := f.Type.Align()
			//fsize := int(f.Type.Size())
			foffset := offset + int(f.Offset)
			//fmt.Printf("T: %-10sS: %d\tA: %d\tO: %d\n", f.Type.Kind(), fsize, falign, foffset)
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
			g[y][x] = 'x'
			x++
		}
	}
}

func (g Layout) String() string {
	if len(g) < 1 {
		return ""
	}
	var buf bytes.Buffer
	for i, line := range g {
		if i > 0 {
			buf.WriteByte('\n')
		}
		buf.Write(line)
	}
	return buf.String()
}
