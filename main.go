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
	allMonth := months[4:]
	fmt.Printf("this is all the month: %v", allMonth)
}
