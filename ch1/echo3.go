package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Running command: " + os.Args[0])
	fmt.Println(strings.Join(os.Args[:], " "))
}

