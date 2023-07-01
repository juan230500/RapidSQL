package main

import "fmt"

func test1() {
	fmt.Println("test1")
}

func test2() {
	fmt.Println("test2-1")

	defer test1()

	return

	fmt.Println("test2-2")
}

func main() {
	test2()
	fmt.Println("main")
}
