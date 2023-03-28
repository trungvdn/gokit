package main

import "fmt"

func main() {

	a := &A{
		name: "a",
		list: []B{
			{
				link: "linkA",
			},
			{
				link: "linkB",
			},
		},
	}
	listA := []A{}
	for i := 0; i < 5; i++ {
		copyA := *a
		copyA.name = "b"
		if i == 3 {
			for _, v := range copyA.list {
				v.link = "linkC"
			}
		}
		for _, v := range copyA.list {
			fmt.Println("link ", v.link)
		}
		listA = append(listA, copyA)
	}
	for _, v := range listA {
		fmt.Println("name = ", v.name)
		for _, va := range v.list {
			fmt.Println("list = ", va.link)
		}
	}
}

type A struct {
	name string
	list []B
}

type B struct {
	link string
}
