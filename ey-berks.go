package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"github.com/ukitazume/ey-berks/config"
	"os"
)

func main() {
	os.Exit(Command(os.Args[1:]))
}

func Command(argv []string) int {
	usage := `Engine Yard Cloud cookbook berkshelf

Usage:
  ey-berks init <path>
	ey-berks make <path>
	ey-berks help
	ey-berks -v | --version
`

	args, _ := docopt.Parse(usage, argv, true, "", false)

	if args["help"] == true {
		fmt.Print(usage)
		return 0
	} else if args["init"] == true {
		fmt.Printf("create %s at %s\n", config.ConfigFileName, args["<path>"])
		if err := config.Create(argv[1]); err != nil {
			fmt.Println(err)
			return 1
		}
	} else if args["--version"] == true || args["-v"] == true {
		fmt.Printf("ey-berks version is: %s", "0.1")
	} else {
		fmt.Println("The command doesn't exist.Please check ey-berks help.")
		return 0
	}
	return 0
}
