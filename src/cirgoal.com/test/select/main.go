package main

import "fmt"

func main() {
	a := make(chan string,10)
	b := make(chan  int, 10)

	for i := 0 ; i< 10 ; i++ {
		a <- "A"
	}

	for i := 0 ; i< 10 ; i++ {
		b <- i
	}

	for  {
		select {
		case <- a:
			fmt.Println("A")
		case <- b :
			fmt.Println("B")
		}
	}

}
