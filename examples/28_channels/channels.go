package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() { 
		messages <- "ping"
		messages <- "pong"
	}()
	
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println("done")
}