package main

import(
	"fmt"
)

// It not that efficient, we can crete a group of consts
const (
	language = "Go"
	maxRetries = 3
)

var(
	pendingCount int
	title string = "Values"
	enabled = true
)

func main() {
	localCount := 2
	// we can declare multiple variables at the same time with GO
	limit, label := 3,"tasks"

	fmt.Println("language: ", language)
	fmt.Println("maxRetries: ",maxRetries)
	fmt.Println("pendingCount: ",pendingCount)
	fmt.Println("title; ",title)
	fmt.Println("enabled; ",enabled)
	fmt.Println("localCount; ",localCount)
	fmt.Println("limit; ",limit)
	fmt.Println("label; ",label)

	value := 42
	//var pointer *int = &value
	pointer := &value
	fmt.Println(value)
	fmt.Println(*pointer) // indirection: *pointer just read what is inside that address

	*pointer = 32 // here is mutation, that changes what is in *pointer address
	fmt.Println(value)
	fmt.Println(*pointer) 

	role := "pasivo"
	roleP := &role
	fmt.Println(role)
	fmt.Println(roleP)
	fmt.Println(*roleP) //*pointer reads what is in the address

	*roleP = "Top"
	fmt.Println(*roleP)
	fmt.Println(role)
}