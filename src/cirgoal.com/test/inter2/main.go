package main

import (
	"fmt"
	"math"
)

func main() {
	c := &Circle{10.0}
	s := &Square{10.0}
	measure(c)
	measure(s)
	c.size(5.0)
	s.size(5.0)
	measure(c)
	measure(s)
}

type Shape interface {
	area() float64
	length() float64
}

type Circle struct {
	redis float64
}

type Square struct {
	edge float64
}

type Change interface {
	size(s float64)
}

func (c *Circle) size(s float64)  {
	c.redis = s
}

func (c *Circle)area() float64 {
	return c.redis*c.redis *math.Pi
}

func (c *Circle) length() float64 {
	return c.redis*2*math.Phi
}

func (s *Square) size(size float64)  {
	s.edge = size
}

func (s *Square) area() float64 {
	return s.edge*s.edge
}

func (s *Square) length() float64 {
	return s.edge*4
}

func measure(s Shape)  {
	fmt.Println(s.area())
	fmt.Println(s.length())
}