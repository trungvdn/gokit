package main

import (
	"fmt"
)

type Learner struct {
	lessonId string
}

type LessonSearch struct {
	lessonId string
	learners []Learner
}

func main() {
	lessonDoc := &LessonSearch{
		lessonId: "a",
	}
	learners := []Learner{}
	lessonDoc.learners = learners
	fmt.Println(lessonDoc)
	fmt.Println(lessonDoc.learners)

}
