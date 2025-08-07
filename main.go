package main

import (
	"fmt"
	"os"
	"flag"
)



func main() {
	fmt.Println("args: ", os.Args[1:])
	method := flag.String("X", "GET", "HTTP method (GET, POST, etc.)")
	flag.Parse()
	fmt.Println(*method)	
}
