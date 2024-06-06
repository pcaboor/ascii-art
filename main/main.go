package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	//security
	if len(os.Args) != 2 {
		fmt.Println("Please provide only one .txt file :)")
		return
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	chr := strings.Split(string(data), "")

	fmt.Println(chr)

}
