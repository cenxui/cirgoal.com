package main

import (
	"fmt"
	"sync"
)

func main() {

	var s sync.WaitGroup
	s.Add(10)
	for i := 0 ;i < 10 ; i++ {
		go run(i, &s)
	}
	s.Wait()
	fmt.Println("Complete")
}

func run(index int, s *sync.WaitGroup)  {
	fmt.Printf("run %d\n", index)
	defer s.Done()
}