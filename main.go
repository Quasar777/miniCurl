package main

import (
	"flag"
	"fmt"
	"strings"
)



func main() {
	method := flag.String("X", "GET", "HTTP method")
	headers := flag.String("H", "", "Headers of request")
	body := flag.String("d", "", "Body of request")
	fileNameForSaving := flag.String("o", "", "file name for saving")
	flag.Parse()

	fmt.Println("method: ", *method)	
	GetHeaders(*headers)
	fmt.Println("body: ", *body)	
	fmt.Println("file for saving: ", *fileNameForSaving)	

	GetHeaders(*headers)
}

func GetHeaders(rawStr string) []string {
	headers := strings.Split(rawStr, ";")

	/*
	for i := range headers {
		headers[i] = strings.Trim(headers[i], " ")
		fmt.Println(headers[i])
	}
	*/

	return headers
}
