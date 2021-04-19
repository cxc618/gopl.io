package main

import (
	"fmt"
	"gopl.io/my/basic/greetings"
	"log"
	"os"
)

func main()  {
	//message, err := greetings.Hello(os.Args[1])
	names := []string{"cxc", "hh", "zz"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)

	if len(os.Args) == 2 {
	} else {
		fmt.Printf("error args %s\n", os.Args)
	}
}
