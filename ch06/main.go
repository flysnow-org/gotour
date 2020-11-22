package main

import (
	"fmt"
)

func main() {
	p := person{
		age:  30,
		name: "飞雪无情",
		address: address{
			province: "北京",
			city:     "北京",
		},
	}
	fmt.Println(p.name, p.age)
	fmt.Println(p.province)
	printString(p.address)
	printString(&p)

	p1 := NewPerson("张三")
	fmt.Println(p1)

	var s fmt.Stringer
	s = p1
	p2 := s.(*person)
	fmt.Println(p2)
	a, ok := s.(address)
	if ok {
		fmt.Println(a)
	} else {
		fmt.Println("s不是一个address")
	}

	add := address{province: "北京", city: "北京"}
	printString(add)
	printString(&add)
}

type person struct {
	name string
	age  uint

	address
}

func NewPerson(name string) *person {
	return &person{name: name}
}

func (p *person) String() string {
	return fmt.Sprintf("the name is %s,age is %d", p.name, p.age)
}

func (p *person) Walk() {
	fmt.Printf("%s能走\n", p.name)
}

func (p *person) Run() {
	fmt.Printf("%s能跑\n", p.name)
}

func printString(s fmt.Stringer) {
	fmt.Println(s.String())
}

type address struct {
	province string
	city     string
}

func (addr address) String() string {
	return fmt.Sprintf("the addr is %s%s", addr.province, addr.city)
}

type WalkRun interface {
	Walk()
	Run()
}
