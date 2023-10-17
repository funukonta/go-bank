package main

import "fmt"

func main() {
	fmt.Println("Yeah")
	server := NewAPIServer(":5000")

	server.Run()
}
