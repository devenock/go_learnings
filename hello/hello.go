package hello

import "fmt"

func Hello(name string) {
	fmt.Println("What's your name?")
	fmt.Scan(&name)
	fmt.Printf("Hello, %v!\n", name)
}
