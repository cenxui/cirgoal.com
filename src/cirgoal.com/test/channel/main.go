package main

import "fmt"

func main() {

	natural := make(chan int)
	squiral := make(chan int)

	go func() {
		for i:= 1; ; i++ {
			natural <- i
			fmt.Printf("natural %d \n", i)
			if i >100 {
				close(natural)
				break
			}
		}

	}()

	go func() {
		for true {
			x, ok := <- natural
			if !ok {
				break
			}

			squiral <- x * x
		}
	}()

	for true {
		v := <-squiral
		fmt.Printf("squiral %d \n", v)
	}

	close(squiral)
}
