package util

import (
	"fmt"
	"reflect"
	"testing"
)

//func TestGetAllCourse(t *testing.T) {
//	GetAllCourse()
//}
//
//func TestGetCourseMenu(t *testing.T)  {
//	GetCourseMenu(605)
//}
//
//func TestGetCourseContent(t *testing.T) {
//	GetCourseContent(6301)
//}

func TestRun(t *testing.T) {
	Run()
}

func TestPaincRevover(t *testing.T) {
	type person struct {
		age  int
		name string
	}

	p := person{12, "Tom"}

	fmt.Println(p)
	var s *string = &p.name

	*s = "jack"
	fmt.Println(p)
}

func TestReflect(t *testing.T) {
	i := 3

	iv := reflect.ValueOf(i)
	it := reflect.TypeOf(i)

	if ii, ok := it.(reflect.Type); ok {
		fmt.Println(ii)
	}

	i1 := iv.Interface()
	i1 = 323
	fmt.Println(i1)

	fmt.Println(iv, it)
}

func TestJsonStr(t *testing.T) {
	type person struct {
		Age  int    `json:"age"`
		Name string `json:"name"`
	}

	var per = person{Age: 18, Name: "hyj"}

	pv := reflect.ValueOf(per)
	pt := reflect.TypeOf(per)
	nums := pt.NumField()

	for i := 0; i < nums; i++ {
		fmt.Println(pt.Field(i).Tag.Get("json"))
		fmt.Println(pt.Field(i).Type)
		fmt.Println(pv.Field(i))
	}
}
