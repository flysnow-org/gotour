package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type MyInt int

func main() {
	fmt.Println(minInt(2, 3))
	fmt.Println(minFloat32(1.2, 2.4))

	fmt.Println(min(1, 2))
	fmt.Println(min(1.2, 2.4))
	fmt.Println(min(1, 2))
	fmt.Println(min("飞雪", "无情"))

	s := []int{1, 2, 3, 4, 5}
	fmt.Println(map1(s, func(v int) int {
		return v * 2
	}))

	s1 := MySlice{1, 3, 5, 7, 9}
	fmt.Println(s1.map1(func(v int) int {
		return v * 2
	}))

	gs1 := GSlice[int]{1, 3, 5, 7, 9}
	fmt.Println(gs1.map1(func(v int) int {
		return v * 2
	}))

	gss1 := GSlice[string]{"1", "3", "5", "7", "9"}
	fmt.Println(gss1.map1(func(v string) string {
		return v + "飞雪"
	}))

	sum := gs1.map1(func(v int) int {
		return v * 2
	}).reduce(func(preV int, curV int) int {
		return preV + curV
	})
	fmt.Println(sum)

	fmt.Println(gs1.filter(func(v int) bool {
		return v > 3
	}))

	sum1 := gs1.filter(func(v int) bool {
		return v > 3
	}).map1(func(v int) int {
		return v * 2
	}).reduce(func(preV int, curV int) int {
		return preV + curV
	})
	fmt.Println("filter map reduce >> sum:", sum1)
}

func minInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func minFloat32(a, b float32) float32 {
	if a < b {
		return a
	} else {
		return b
	}
}

func min[V constraints.Ordered](a, b V) V {
	if a < b {
		return a
	} else {
		return b
	}
}

type Number interface {
	int | float32 | float64 | int32
}

func map1(s []int, f func(int) int) []int {
	result := make([]int, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

type MySlice []int

func (s MySlice) map1(f func(int) int) MySlice {
	result := make(MySlice, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

type GSlice[T any] []T

func (s GSlice[T]) map1(f func(T) T) GSlice[T] {
	result := make(GSlice[T], len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

func (s GSlice[T]) reduce(f func(previousValue T, currentValue T) T) T {
	var result T
	for _, v := range s {
		result = f(result, v)
	}
	return result
}

func (s GSlice[T]) filter(f func(T) bool) GSlice[T] {
	result := GSlice[T]{}
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}
