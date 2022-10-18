package main

import (
	"fmt"

	"github.com/vsantos1/websockets-go/route"
)



func main() {
	route.NewRouter()
	fmt.Println("Hello, World!")
}