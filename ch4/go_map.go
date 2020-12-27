package main

import "fmt"

func main() {
	// 1. first method to init map
	ages := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34

	delete(ages, "alice")

	ages["charlie"] = ages["crise"] + 1
	ages["charlie"] += 1
	ages["charlie"]++
	ages["crise"] = 26
	// 2. second method to init map
	//ages := map[string]int {
	//	"alice": 31,
	//	"charlie": 34,
	//}
	fmt.Println(ages)

	// 3. traverse map
	for key, value := range ages {
		fmt.Printf("%s --> %d\n", key, value)
	}

	fmt.Println(len(ages))
	ages["zhao"] = 20
}
