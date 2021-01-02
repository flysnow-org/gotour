package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	ss := []string{"飞雪无情", "张三"}
	fmt.Println("切片ss长度为", len(ss), ",容量为", cap(ss))
	ss = append(ss, "李四", "王五")
	fmt.Println("切片ss长度为", len(ss), ",容量为", cap(ss))
	fmt.Println(ss)

	a1 := [2]string{"飞雪无情", "张三"}
	fmt.Printf("函数main数组指针：%p\n", &a1)
	arrayF(a1)

	s1 := a1[0:1]
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s1)).Data)
	sliceF(s1)

	s2 := a1[:]
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s2)).Data)

	sh1 := (*slice)(unsafe.Pointer(&s1))
	fmt.Println(sh1.Data, sh1.Len, sh1.Cap)

	s := "飞雪无情"
	fmt.Printf("s的内存地址：%d\n", (*reflect.StringHeader)(unsafe.Pointer(&s)).Data)
	b := []byte(s)
	fmt.Printf("b的内存地址：%d\n", (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data)
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	sh.Cap = sh.Len
	b1 := *(*[]byte)(unsafe.Pointer(sh))
	fmt.Printf("b1的内存地址：%d\n", (*reflect.SliceHeader)(unsafe.Pointer(&b1)).Data)
	s3 := string(b)
	fmt.Printf("s3的内存地址：%d\n", (*reflect.StringHeader)(unsafe.Pointer(&s3)).Data)
	s4 := *(*string)(unsafe.Pointer(&b))
	fmt.Printf("s4的内存地址：%d\n", (*reflect.StringHeader)(unsafe.Pointer(&s4)).Data)
	fmt.Println(s, string(b), string(b1), s3)
}

func arrayF(a [2]string) {
	fmt.Printf("函数arrayF数组指针：%p\n", &a)
}

func sliceF(s []string) {
	fmt.Printf("函数sliceF Data：%d\n", (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data)
}

type slice struct {
	Data uintptr
	Len  int
	Cap  int
}
