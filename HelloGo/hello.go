package main

import "fmt"

func main() {
	HelloGo("Go")
}

func HelloGo(name string) {
	fmt.Println(fmt.Sprintf("Hello, %s!", name))
}
