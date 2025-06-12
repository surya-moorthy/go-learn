package main

import "fmt"

func main() {
	go hello()
	fmt.Println("The main print function")
}

func hello() {
	fmt.Println("The is the first time i am using go routine")
}

// here the main function runs before the hello function get printed , 
// so the main thing here we need to channel this out how the flow of the program would or wait until the hello function gets printed. 