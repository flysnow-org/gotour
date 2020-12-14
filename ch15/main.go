package main

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strings"
)

func main() {
	i:=3
	iv:=reflect.ValueOf(i)
	it:=reflect.TypeOf(i)
	fmt.Println(iv.Int(),it)//3,int

	i1:=iv.Interface().(int)
	fmt.Println(i1)

	ipv:=reflect.ValueOf(&i)
	ipv.Elem().SetInt(4)
	fmt.Println(i)

	p:=person{Name: "飞雪无情",Age: 20}
	ppv:=reflect.ValueOf(&p)
	ppv.Elem().Field(0).SetString("张三")
	fmt.Println(p)
	fmt.Println(ppv.Kind())

	pv:=reflect.ValueOf(p)
	fmt.Println(pv.Kind())

	pt:=reflect.TypeOf(p)
	//遍历person的字段
	for i:=0;i<pt.NumField();i++{
		fmt.Println("字段：",pt.Field(i).Name)
	}
	//遍历person的方法
	for i:=0;i<pt.NumMethod();i++{
		fmt.Println("方法：",pt.Method(i).Name)
	}

	stringerType:=reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	writerType:=reflect.TypeOf((*io.Writer)(nil)).Elem()
	fmt.Println("是否实现了fmt.Stringer：",pt.Implements(stringerType))
	fmt.Println("是否实现了io.Writer：",pt.Implements(writerType))

	//struct to json
	jsonB,err:=json.Marshal(p)
	if err==nil {
		fmt.Println(string(jsonB))
	}

	//json to struct
	respJSON:="{\"name\":\"李四\",\"age\":40}"
	json.Unmarshal([]byte(respJSON),&p)
	fmt.Println(p)

	//遍历person字段中key为json、bson的tag
	for i:=0;i<pt.NumField();i++{
		sf:=pt.Field(i)
		fmt.Printf("字段%s上,json tag为%s\n",sf.Name,sf.Tag.Get("json"))
		fmt.Printf("字段%s上,bson tag为%s\n",sf.Name,sf.Tag.Get("bson"))
	}

	jsonBuilder:=strings.Builder{}
	jsonBuilder.WriteString("{")
	num:=pt.NumField()
	for i:=0;i<num;i++{
		jsonTag:=pt.Field(i).Tag.Get("json")
		jsonBuilder.WriteString("\""+jsonTag+"\"")
		jsonBuilder.WriteString(":")
		jsonBuilder.WriteString(fmt.Sprintf("\"%v\"",pv.Field(i)))
		if i<num-1{
			jsonBuilder.WriteString(",")
		}
	}
	jsonBuilder.WriteString("}")
	fmt.Println(jsonBuilder.String())

	//反射调用person的Print方法
	mPrint:=pv.MethodByName("Print")
	args:=[]reflect.Value{reflect.ValueOf("登录")}
	mPrint.Call(args)
}

type person struct {
	Name string `json:"name" bson:"b_name"`
	Age int `json:"age" bson:"b_name"`
}

func (p person) String() string{
	return fmt.Sprintf("Name is %s,Age is %d",p.Name,p.Age)
}

func (p person) Print(prefix string){
	fmt.Printf("%s:Name is %s,Age is %d\n",prefix,p.Name,p.Age)
}