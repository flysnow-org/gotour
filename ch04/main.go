package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	array:=[5]string{"a","b","c","d","e"}
	fmt.Println(array[2])
	array1:=[5]string{1:"b",3:"d"}
	fmt.Println(array1)

	for i:=0;i<5;i++{
		fmt.Printf("数组索引:%d,对应值:%s\n", i, array[i])
	}

	for i,v:=range array{
		fmt.Printf("数组索引:%d,对应值:%s\n", i, v)
	}

	slice:=array[2:5]
	slice[1] ="f"
	fmt.Println(array)

	slice1:=[]string{"a","b","c","d","e"}

	slice2:=append(slice1,"f")
	fmt.Println(slice1,slice2)

	nameAgeMap:=make(map[string]int)
	nameAgeMap["飞雪无情"] = 20

	age,ok:=nameAgeMap["飞雪无情1"]
	if ok {
		fmt.Println(age)
	}
	delete(nameAgeMap,"飞雪无情")

	//测试for range
	nameAgeMap["飞雪无情"] = 20
	nameAgeMap["飞雪无情1"] = 21
	nameAgeMap["飞雪无情2"] = 22

	for k,v:=range nameAgeMap{
		fmt.Println("Key is",k,",Value is",v)
	}

	fmt.Println(len(nameAgeMap))

	s:="Hello飞雪无情"
	bs:=[]byte(s)
	fmt.Println(bs)
	fmt.Println(s[0],s[1],s[15])
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	for i,r:=range s{
		fmt.Println(i,r)
	}

	aa:=[3][3]int{}
	aa[0][0] =1
	aa[0][1] =2
	aa[0][2] =3
	aa[1][0] =4
	aa[1][1] =5
	aa[1][2] =6
	aa[2][0] =7
	aa[2][1] =8
	aa[2][2] =9
	fmt.Println(aa)
}
