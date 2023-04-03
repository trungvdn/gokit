package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
)

// before
// type circle struct {
// 	radius float64
// }

// func (c *circle) area() {
// 	fmt.Printf("circle area: %f\n", math.Pi*c.radius*c.radius)
// }

// func main() {

// 	c := circle{
// 		radius: 3,
// 	}
// 	c.area()

// }

// refactor

type shape interface {
	name() string
	area() float64
}

type circle struct {
	radius float64
}

func (c *circle) name() string {
	return "circle"
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type square struct {
	length float64
}

func (s *square) name() string {
	return "square"
}

func (s *square) area() float64 {
	return s.length * s.length
}

type output struct{}

func (o *output) Text(sh shape) string {
	return fmt.Sprintf("area of shape %s is %f", sh.name(), sh.area())
}

func (o *output) JSON(sh shape) string {
	res := struct {
		Name string  `json:"shape"`
		Area float64 `json:"area"`
	}{
		Name: sh.name(),
		Area: sh.area(),
	}

	bs, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}
	return string(bs)
}
