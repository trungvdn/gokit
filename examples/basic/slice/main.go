package main

import "fmt"

type Learner struct {
	lessonId string
}

type LessonSearch struct {
	lessonId string
	learners []Learner
}

func main() {
	// lessonDoc := &LessonSearch{
	// 	lessonId: "a",
	// }
	// learners := []Learner{}
	// lessonDoc.learners = learners
	// fmt.Println(lessonDoc)
	// fmt.Println(lessonDoc.learners)

	// var s []int
	// fmt.Println(s == nil) // True because slice just a reference to an array
	// fmt.Printf("%v\n", s)
	// fmt.Println("len", len(s))
	// fmt.Println("cap", cap(s))

	// var s []int
	// arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8}

	// s = arr[1:3]
	// fmt.Println("len,cap", len(s), cap(s))

	// arr[0] = 11
	// arr[2] = 22
	// fmt.Println(s)
	// fmt.Println(arr)
	// s[0] = 90
	// s[1] = 13
	// fmt.Println(s)
	// fmt.Println(arr)

	// fmt.Println("address of arr[2]", &arr[2])
	// fmt.Println("address of s[1]", &s[1])

	// a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// s := a[2:4]
	// fmt.Printf("before -> s=%v\n", s)
	// fmt.Printf("before -> a=%v\n", a)
	// fmt.Printf("before -> len=%d, cap=%d\n", len(s), cap(s))
	// fmt.Println("&a[2] == &s[0] is", &a[2] == &s[0])

	// s = append(s, 50, 60, 70, 80, 90, 100)
	// // s = append(s, 50, 60, 70, 80, 90, 100, 110)
	// fmt.Printf("after -> s=%v\n", s)
	// fmt.Printf("after -> a=%v\n", a)
	// fmt.Printf("after -> len=%d, cap=%d\n", len(s), cap(s))
	// fmt.Println("&a[2] == &s[0] is", &a[2] == &s[0])

	countries := getCountries()
	fmt.Printf("after -> len=%d, cap=%d\n", len(countries), cap(countries))
	fmt.Println(countries) // 9
	fmt.Println(countries[4])

}

func getCountries() []string {
	countries := []string{"United states", "United kingdom", "Austrilia", "India", "China", "Russia", "France", "Germany", "Spain"} // can be much more
	return countries[:3]
}
