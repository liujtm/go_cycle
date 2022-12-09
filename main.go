package main

import "fmt"

func main() {
	fmt.Println("hello world")
	app, _ := InitializeApplication()
	app.A.PrintAll()
	app.B.PrintAll()
}
