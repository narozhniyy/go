package main

import "fmt"

func main() {
	ch := make(chan int)
	q := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		q <- 999
	}()
	for {
		select {
		case a := <-ch:
			fmt.Println(a)
		case quit := <-q:
			fmt.Println(quit)
			return
		}
	}
}
