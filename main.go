package main

import "fmt"

func main() {
	fmt.Println("hello world")
	app, _ := InitializeApplication()

	app.A.PrintA()
	app.A.PrintAll()

	app.B.PrintB()
	app.B.PrintAll()
}
