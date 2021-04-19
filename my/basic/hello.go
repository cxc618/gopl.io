package main

import (
	"fmt"
	"os"
)

func main()  {
	len := len(os.Args)
	if (len == 2) {
		fmt.Printf("Hello, %s\n", os.Args[1])
	} else {
		fmt.Printf("error args %s\n", os.Args)
	}
}
