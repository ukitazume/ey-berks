package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	author "github.com/ukitazume/ey-berks/author"
	"github.com/ukitazume/ey-berks/config"
	"github.com/ukitazume/ey-berks/gather"
	"os"
)

const (
	Version = "0.0.1"
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
		path := args["<path>"].(string)
		if err := config.Create(path); err != nil {
			fmt.Println(err)
			return 1
		}
	} else if args["make"] == true {
		path := args["<path>"].(string)

		g := gather.NewGather(path)
		if err := g.Gather(path); err != nil {
			fmt.Println(err)
		}

		list := author.CreateMainRecipe(g.Berks)
		if err := author.CreateFile(path, list); err != nil {
			fmt.Printf("error: %v\\\\n", err)
		}
	} else if args["--version"] == true || args["-v"] == true {
		fmt.Printf("ey-berks version is: %s", Version)
	} else {
		fmt.Println("The command doesn't exist.Please check ey-berks help.")
		return 0
	}
	return 0
}
