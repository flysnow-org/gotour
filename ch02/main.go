package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var i int = 10
	var f32 float32 = 2.2
	var f64 float64 = 10.3456
	var bf bool =false
	var bt bool = true
	var s1 string = "Hello"
	var s2 string = "世界"
	s1+=s2
	fmt.Println(i)
	fmt.Println("f32 is",f32,",f64 is",f64)
	fmt.Println("bf is",bf,",bt is",bt)
	fmt.Println("s1 is",s1,",s2 is",s2)
	fmt.Println("s1+s2=",s1+s2)

	var zi int
	var zf float64
	var zb bool
	var zs string
	fmt.Println(zi,zf,zb,zs)

	pi:=&i
	fmt.Println(*pi)

	i = 20
	fmt.Println("i的新值是",i)
	const name =  "飞雪无情"

	const(
		one = iota+1
		two
		three
		four
	)
	fmt.Println(one,two,three,four)

	i2s:=strconv.Itoa(i)
	s2i,err:=strconv.Atoi(i2s)
	fmt.Println(i2s,s2i,err)

	i2f:=float64(i)
	f2i:=int(f64)
	fmt.Println(i2f,f2i)

	//判断s1的前缀是否是H
	fmt.Println(strings.HasPrefix(s1,"H"))
	//在s1中查找字符串o
	fmt.Println(strings.Index(s1,"o"))
	//把s1全部转为大写
	fmt.Println(strings.ToUpper(s1))
}
