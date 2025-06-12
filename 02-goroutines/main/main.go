package main

import "fmt"

// fatal error: all goroutines are asleep - deadlock!

// this error happens because there is no outside interaction between go routine and channel inside the main function

// if there is a sender to send then there must be a receiver to receive.

func main() {
   ch := make(chan int)

   ch <- 10

   vh := <- ch
   fmt.Println("received",vh)
}