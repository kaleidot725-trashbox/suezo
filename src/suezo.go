package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	source := flag.String("src", "./", "help message for src")
	destination := flag.String("dst", "./", "help message for dst")
	deleteOption := flag.Bool("delete", false, "help message for d option")
	helpOption := flag.Bool("help", false, "help message for h option")
	versionOption := flag.Bool("version", false, "help message for v option")
	flag.Parse()

	fmt.Printf("sss %t\n ", *deleteOption)

	if 3 <= len(os.Args) {
		organize(*source, *destination, *deleteOption)
		return
	}

	if *helpOption {
		menu()
		return
	}

	if *versionOption {
		version()
		return
	}

	menu()
}

func menu() {
	fmt.Println("suezo is a tool to forced stacks")
	fmt.Println()
	fmt.Println("Usage :")
	fmt.Println("     suezo <source> <destination> [options]")
	fmt.Println()
	fmt.Println("The options are")
	fmt.Println("     -d     delete source directories.")
	fmt.Println("     -h     display help message.")
	fmt.Println("     -v     display version message.")
	fmt.Println()
}

func version() {
	fmt.Println("Suezo is v0.0.1")
}

func organize(source string, destination string, remove bool) {
	organizer := Organizer{Explorer{}}
	err := organizer.OriganizeByExtension(source, destination, remove)
	if err != nil {
		fmt.Printf("organization failed!!(%s)", err)
		return
	}

	fmt.Println("organization success!!")
	return
}
