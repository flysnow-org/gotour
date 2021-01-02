package main

import (
	"fmt"
	"os"
)

const name = "飞雪无情"

func main() {
	os.Mkdir("tmp", 0666)

	fmt.Println("飞雪无情")

	m := map[int]string{}
	s := "飞雪无情"
	m[0] = s
}

func newString() string {
	s := new(string)
	*s = "飞雪无情"
	return *s
}
