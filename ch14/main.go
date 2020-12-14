package main

import "fmt"

func main() {
	var s string
	fmt.Printf("%p\n", &s)
	s = "张三"
	fmt.Println(s)

	var sp *string
	sp = new(string)
	*sp = "飞雪无情"
	fmt.Println(*sp)

	pp:=NewPerson("飞雪无情",20)
	fmt.Println("name为",pp.name,",age为",pp.age)

	m:= map[string]int{"张三":18}
	fmt.Println(m)

}

func NewPerson(name string,age int) *person{
	p:=new(person)
	p.name = name
	p.age = age
	return p
}

type person struct {
	name string
	age int
}