package main

import (
	"fmt"
	"strconv"

	"github.com/modeckrus/utils"
)

type Person struct {
	Age      int
	Name     string
	Surname  string
	Children []Person
}

func main() {
	for i := 1; i < 100; i++ {
		children := make([]Person, i)
		for g := 0; g < i; g++ {
			children[g] = Person{
				Name:    strconv.Itoa(g),
				Surname: strconv.Itoa(g),
				Age:     g,
			}
		}
		fmt.Print(utils.SPrint(Person{
			Name:     strconv.Itoa(i),
			Surname:  strconv.Itoa(i),
			Age:      i,
			Children: children,
		}) + "\n")
	}
	// utils.SPrintf("Ok: %v", 12)
}
