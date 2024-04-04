package main

import "fmt"

func main() {
	for i := 0; i < 40; i++ {
		go GetInstance()
	}

	fmt.Scanln()
}
