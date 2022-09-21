package main

import (
	"fmt"
)

func printMyGithubName(githubName string) {
	fmt.Println(githubName)
}

// example of pointers
//func pointer() {
//	var four int = 4
//	addressOfFour := &four
//	fmt.Printf("this is a pointer of: %v", addressOfFour)
//}

type naira float32

type kobo float32

func (k kobo) naira() naira {
	return naira(k / 100)
}

func main() {
	kobo := kobo(100)
	fmt.Println(kobo.naira())
}

func _main() {
	printMyGithubName("hussain4real\n")

	var months []string

	/*var songName = []string{
		"song1",
		"song2",
		"song3",
	}*/

	months = []string{
		"jan",
		"feb",
		"mar",
		"apr",
		"may",
		"jun",
		"jul",
		"aug",
		"sep",
		"oct",
		"nov",
		"dec",
	}
	allMonth := months[:]

	fmt.Printf("this is all the month: %v\n", allMonth)

	allMonth = append(allMonth, "hello")
	fmt.Println(allMonth)

	//deep copy example
	newMonth := make([]string, len(allMonth))
	copy(newMonth, allMonth)
	fmt.Println(newMonth)

	//shallow copy example
	items := []string{"phones", "chargers"}
	copyItems := items

	fmt.Printf("address of items is: %[1]p \n", items)
	fmt.Printf("address of items is: %[1]p \n", copyItems)

	copyItems = append(copyItems, "comb")

	fmt.Println(items)
	fmt.Println(copyItems)

	for c := 0; c <= 10; c++ {
		fmt.Printf("c is:%v \n", c)
	}
}
