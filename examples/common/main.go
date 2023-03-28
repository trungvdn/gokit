package main

import (
	"fmt"
	"strconv"
	"time"

	"golang.org/x/exp/slices"
)

type Color uint8

const (
	Red Color = iota
	Blue
)

func main() {
	alist := []Apple{
		{
			name:  "T1",
			color: Blue,
		},
		{
			name:  "T2",
			color: Red,
		},
		{
			name:  "T3",
			color: Blue,
		},
	}

	res := filter(alist, func(a Apple) bool {
		return a.color == Blue
	})
	fmt.Println(res)
	// s1 := []string{"1", "2"}
	// s2 := []string{"1", "2"}
	fmt.Println(slices.Equal([]string{"1"}, []string{"11"}))
	s1 := []string{"1", "2"}
	s2 := []int{1, 3}
	fmt.Println(slices.EqualFunc(s1, s2, func(e1 string, e2 int) bool {
		return e1 == strconv.Itoa(e2)
	}))
	fmt.Println(slices.Compact([]string{"1", "1", "3", "3", "3", "4", "4"}))
	fmt.Println(mapForEach([]string{"1", "2"}, func(s string) int {
		i, _ := strconv.Atoi(s)
		return i
	}))
}

type ApplePredicate struct {
}

func (a *ApplePredicate) test(p Apple) bool {
	return p.color == Blue
}

type Apple struct {
	name  string
	color Color
}

type Predicate interface {
	test(Apple) bool
}

func anyMatch[T any](list []T, fn func(T) bool) bool {
	for _, a := range list {
		if fn(a) {
			return true
		}
	}
	return false
}

func allMatch[T any](list []T, fn func(T) bool) bool {
	for _, a := range list {
		if !fn(a) {
			return false
		}
	}
	return true
}

func filter[T any](list []T, fn func(T) bool) []T {
	var res []T
	for _, a := range list {
		if fn(a) {
			res = append(res, a)
		}
	}
	return res
}

func mapForEach(arr []string, fn func(string) int) []int {
	res := []int{}
	for _, v := range arr {
		res = append(res, fn(v))
	}
	return res
}

func consumer(arr []string, fn func(string)) {
	for _, v := range arr {
		fn(v)
	}
}

type Lesson struct {
	id   string
	name string
	date time.Time
}
