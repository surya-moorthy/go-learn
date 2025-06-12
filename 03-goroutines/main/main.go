package main

import "fmt"

func main() {
   ch := make(chan string)
	
	go func() {
	   ch <- "hi"
	}()

    get := <- ch

	fmt.Println("received",get)

	go hello(ch)

   vh := <- ch
   fmt.Println("received",vh)
}


func hello(ch chan string) {
   fmt.Println("giving a hello message")
   ch <- "hello"	
}

// has been said this is how the channels but i noticed that when first the string "hi" and after that "hello" in the hello fucntion 
// , i think there may be a loop running while channeling the values there is no capacity mentioned , where means how much data should i have receive.
// the solution , i figured i use get variable to get the first data that is channelled out and continued the process.