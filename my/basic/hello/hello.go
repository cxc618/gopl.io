package main

import (
	"fmt"
	"gopl.io/my/basic/greetings"
	"log"
	"os"
)

func main()  {
	message, err := greetings.Hello(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(message)

	if len(os.Args) == 2 {
	} else {
		fmt.Printf("error args %s\n", os.Args)
	}
}
