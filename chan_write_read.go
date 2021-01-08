package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	for j := 0; j < 10; j++ {
		fmt.Println(<-ch)
	}
	close(ch)
	fmt.Println("Hello, playground")
}
