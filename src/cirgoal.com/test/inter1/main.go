package main

import "fmt"

func main() {
	Hello(true)
	Hello("String")
	Hello(10)
	Hello(10.0)

	var x float32
	x = 10
	Hello(x)

}

func Hello(value interface{}) {
	switch value.(type) {
	case string:
		fmt.Println("String")
	case int:
		fmt.Println("Int")
	case bool:
		fmt.Println("Boolean")
	case float32:
		fmt.Println("Float32")
	case float64:
		fmt.Println("Float64")
	}
}

