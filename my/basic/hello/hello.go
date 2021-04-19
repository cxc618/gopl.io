package main

import (
	"fmt"
	"gopl.io/my/basic/greetings"
	"os"
)

func main()  {
	if len(os.Args) == 2 {
		message := greetings.Hello(os.Args[1])
		fmt.Printf(message)
	} else {
		fmt.Printf("error args %s\n", os.Args)
	}
}
