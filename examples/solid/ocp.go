package main

import "fmt"

type Apple struct {
	name   string
	color  string
	weigth int
}

func filterColor(listApple []Apple, color string, weight int) []Apple {
	filteredApple := []Apple{}
	for i := 0; i < len(listApple); i++ {
		if listApple[i].color == color && listApple[i].weigth >= weight {
			filteredApple = append(filteredApple, listApple[i])
		}
	}
	return filteredApple
}

// type Filter interface {
// 	filter(apple Apple) bool
// }

func filterColor2(listApple []Apple, filter Filter) []Apple {
	filteredApple := []Apple{}
	for i := 0; i < len(listApple); i++ {
		if filter(listApple[i]) {
			filteredApple = append(filteredApple, listApple[i])
		}
	}
	return filteredApple
}

type Filter func(apple Apple) bool

type FilterColorAndWeight struct {
	color  string
	weigth int
}

func (fc *FilterColorAndWeight) filter(apple Apple) bool {
	if apple.color == fc.color && apple.weigth >= fc.weigth {
		return true
	}
	return false
}

func ImplementOCP() {
	list := []Apple{
		{
			name:   "Yasuo",
			color:  "Red",
			weigth: 100,
		},
		{
			name:   "Zoe",
			color:  "Pink",
			weigth: 98,
		},
		{
			name:   "Maphite",
			color:  "Red",
			weigth: 144,
		},
	}

	filteredApple := filterColor2(list, func(apple Apple) bool {
		if apple.color == "Red" && apple.weigth >= 100 {
			return true
		}
		return false
	})

	fmt.Println(filteredApple)
}
