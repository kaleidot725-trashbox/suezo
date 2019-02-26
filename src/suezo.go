package main

import (
	"fmt"
)

func main() {
	println("directory")
	explorer := Explorer{}
	for _, file := range explorer.ExploreDirectory("./sample", true) {
		fmt.Println("┗" + file)
	}

	println("file")
	for _, file := range explorer.ExploreFile("./sample", true) {
		fmt.Println("┗" + file)
	}

	return
}
