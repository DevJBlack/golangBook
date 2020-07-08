package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// var s, sep string
	// for i := 1; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i]
	// 	sep = " "
	// }
	// fmt.Println(s)

	// Anothe way to iterate over
	// s, sep := "", ""
	// for _, arg := range os.Args[1:] {
	// 	s += sep + arg
	// 	sep = " "
	// }
	// fmt.Println(s)

	// Third way and the most compact version
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:])
	fmt.Println(os.Args[0])

	for i, v := range os.Args[1:] {
		fmt.Printf(" Index is: %d, Value is: %v", i, v)
	}
}
