package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (lenght int, uppercase string) {
	defer fmt.Println("I'm done")
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return
}


func main() {
	totalLenght, uppercase := lenAndUpper("nico")
	fmt.Println(totalLenght, uppercase)
}