package main

import (
	"fmt"
)

// func main() {

// 	a := &A{
// 		name: "a",
// 		list: []B{
// 			{
// 				link: "linkA",
// 			},
// 			{
// 				link: "linkB",
// 			},
// 		},
// 	}
// 	listA := []A{}
// 	for i := 0; i < 5; i++ {
// 		copyA := *a
// 		copyA.name = "b"
// 		if i == 3 {
// 			for _, v := range copyA.list {
// 				v.link = "linkC"
// 			}
// 		}
// 		for _, v := range copyA.list {
// 			fmt.Println("link ", v.link)
// 		}
// 		listA = append(listA, copyA)
// 	}
// 	for _, v := range listA {
// 		fmt.Println("name = ", v.name)
// 		for _, va := range v.list {
// 			fmt.Println("list = ", va.link)
// 		}
// 	}
// }

type A struct {
	name string
	list []B
}

type B struct {
	link string
}

func changeValue(p *[3]int) {
	(*p)[0] *= 2
	(*p)[1] *= 3
	(*p)[2] *= 4
}

func main() {
	//hexadecimal number
	// a := 0xFF
	// fmt.Printf("%T %v %X", a, a, a)
	// declaring a variable,assign value to it,Go runtime allocate memory in RAM(depend on data type)
	// that memory have address,these memory address represented in hexadecimal value

	a := 1
	pa := &a
	*pa = 2
	fmt.Printf("%v %v\n", pa, *pa)
	fmt.Printf("%v", a)

	// b := 123
	// p := &b
	// fmt.Println("&b = ", p)

	p := new(int)
	*p = 2

	arr := [3]int{1, 2, 3}
	changeValue(&arr)
	fmt.Println(arr)
}
