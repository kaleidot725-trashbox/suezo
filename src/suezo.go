package main

import (
	"flag"
	"os"
)

func main() {
	recursiveOption := flag.Bool("r", false, "help message for r option")
	deleteOption := flag.Bool("d", false, "help message for d option")
	helpOption := flag.Bool("h", false, "help message for h option")
	versionOption := flag.Bool("v", false, "help message for v option")
	flag.Parse()

	if 3 <= len(os.Args) {
		source := os.Args[1]
		destination := os.Args[2]
		organize(source, destination, *recursiveOption, *deleteOption)
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
	println("Suezo is a tool for organizing file")
	println()
	println("Usage :")
	println("     suezo <source> <destination> [options]")
	println()
	println("The options are")
	println("     -r     organize recursively directories.")
	println("     -d     delete source directories.")
	println("     -h     display help message.")
	println("     -v     display version message.")
	println()
}

func version() {
	println("Suezo is v0.0.1")
}

func organize(source string, destination string, recursive bool, remove bool) {
	organizer := Organizer{Explorer{}}
	err := organizer.OriganizeByExtension(source, destination)
	if err != nil {
		println("organization failed!!(%s)", err)
		return
	}

	println("organization success!!")
	return
}
