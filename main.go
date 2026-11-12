package main

import(
	"fmt"
)

// It not that efficient, we can crete a group of consts
const (
	language = "Go"
	maxRetries = 3
)

func main() {
	fmt.Println(language)
	fmt.Println(maxRetries)
}