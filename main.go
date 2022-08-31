package main

import (
	"fmt"
)

func printMyGithubName(githubName string) {
	fmt.Println(githubName)
}

func main() {
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
}
