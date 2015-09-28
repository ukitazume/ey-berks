package main

import (
	"fmt"
	"github.com/ukitazume/ey-berks/config"
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

	switch argv[0] {
	case "help":
		fmt.Print(usage)
		return 0
	case "init":
		fmt.Println("create EYBerksfile")
		err := config.Create(argv[1])
		if err != nil {
			fmt.Print(err)
			return 1
		}
		return 0
	default:
		fmt.Println("The command doesn't exist.Please check ey-berks help.")
		return 0
	}
	return 0
}
