package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"strings"
)

func main() {
	fmt.Println(Stringify([]MyInt{1, 2, 3}))

	p := Point{1, 2, 3}
	ScaleAndPrint(p)
}

type MyInt int

func (i MyInt) String() string {
	return fmt.Sprintf("%d:%d", i, i)
}

func Stringify[T fmt.Stringer](s []T) (ret []string) {
	for _, v := range s {
		ret = append(ret, v.String())
	}
	return ret
}

// Scale returns a copy of s with each element multiplied by c.
func Scale[S ~[]E, E constraints.Integer](s S) S {
	r := make(S, len(s))
	for i, v := range s {
		r[i] = v * 2
	}
	return r
}

type Point []int32

func (p Point) String() string {
	var b strings.Builder
	b.Grow(len(p))
	for _, v := range p {
		b.WriteString(fmt.Sprint(v))
		b.WriteString(",")
	}
	return b.String()
}

// ScaleAndPrint doubles a Point and prints it.
func ScaleAndPrint(p Point) {
	r := Scale(p)
	fmt.Println(r.String())
}
