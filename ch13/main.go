package main

import (
	"fmt"
)

func main() {
	p := person{name: "张三", age: 18}
	fmt.Printf("main函数：p的内存地址为%p\n", &p)
	modifyPerson(&p)
	fmt.Println("person name:", p.name, ",age:", p.age)

	m := make(map[string]int)
	m["飞雪无情"] = 18
	fmt.Println("飞雪无情的年龄为", m["飞雪无情"])
	fmt.Printf("main函数：m的内存地址为%p\n", m)
	modifyMap(m)
	fmt.Println("飞雪无情的年龄为", m["飞雪无情"])

	//类型零值测试
	var s string
	var i int
	var b bool
	var f float64
	var st struct{}
	var mi map[string]int
	var sl []string
	var ia interface{}
	var fn func()
	var ch chan string
	fmt.Println("string的零值为", s)
	fmt.Println("int的零值为", i)
	fmt.Println("bool的零值为", b)
	fmt.Println("float64的零值为", f)
	fmt.Println("struct的零值为", st)
	fmt.Println("map的零值为", mi)
	fmt.Println("slice的零值为", sl)
	fmt.Println("interface的零值为", ia)
	fmt.Println("func的零值为", fn)
	fmt.Println("chan的零值为", ch)
}

func modifyPerson(p *person) {
	fmt.Printf("modifyPerson函数：p的内存地址为%p\n", p)
	p.name = "李四"
	p.age = 20
}

func modifyMap(p map[string]int) {
	fmt.Printf("modifyMap函数：p的内存地址为%p\n", p)
	p["飞雪无情"] = 20
}

type person struct {
	name string
	age  int
}
