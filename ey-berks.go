package main

import (
	"fmt"
	"os"
)

func main() {
	os.Exit(Command(os.Args[1:]))
}

func Command(argv []string) int {

	usage := `
Engine Yard Cloud cookbook berkshelf
Usage: ey-berks <command> [<args>...]

Command:

init     create an EYBerkshelf file
make     make cookbook directory
help     show this usage
`

	if len(argv) == 0 {
		fmt.Print(usage)
		return 0
	}

	fmt.Printf("%s", argv[0])

	switch argv[0] {
	case "help":
		fmt.Print(usage)
		return 0
	default:
		fmt.Print(usage)
		return 0
	}
	return 0
}
