package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	from := flag.String("f", "./", "help message for from")
	to := flag.String("t", "./", "help message for to")
	deleteOption := flag.Bool("delete", false, "help message for d option")
	helpOption := flag.Bool("h", false, "help message for h option")
	versionOption := flag.Bool("v", false, "help message for v option")
	flag.Parse()

	if 5 <= len(os.Args) {
		organize(*from, *to, *deleteOption)
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
	fmt.Println("Suezo is a tool to forced stacks")
	fmt.Println()
	fmt.Println("Usage :")
	fmt.Println("     suezo -src <source> -dst <destination> [options]")
	fmt.Println()
	fmt.Println("The options are")
	fmt.Println("     -delete     delete source directories.")
	fmt.Println("     -help     display help message.")
	fmt.Println("     -version     display version message.")
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
