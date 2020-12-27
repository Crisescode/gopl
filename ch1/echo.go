package main

import (
	"fmt"
	"os"
	"strings"
)

//func main() {
//	var s, sep string
//	for i := 1; i < len(os.Args); i++ {
//		s += sep + os.Args[i]
//		sep = " "
//	}
//
//	fmt.Println(s)
//}

//func main() {
//	fmt.Println(strings.Join(os.Args[1:], " "))
//}

func main() {
	//for idx, value := range os.Args[:] {
	//	fmt.Println(idx, "->", value)
	//}

	fmt.Println(strings.Join(os.Args[:], " "))
}