package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"github.com/ukitazume/ey-berks/author"
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

func usage() string {
	return `Engine Yard Cloud cookbook Berkshelf

Usage: ey-berks <command> [<path>] [--config=<config>]

ey-berks config <path> [--config=<config>]             : make a sample configuration file
ey-berks compile <path> [--config=<config>]            : update cahce,  write a main/recipes and gather recipe to the cookbooks directory
ey-berks update-cache [--config=<config>]              : update cache of remote repositories cookbooks
ey-berks create-main-recipe <path> [--config=<config>]
ey-berks copy-recipes <path> [--config=<config>]
ey-berks help
ey-berks version
`
}

func Command(argv []string) int {

	args, _ := docopt.Parse(usage(), argv, true, "", false)
	command, path, conf := parseArgs(args)

	configOptions := config.DefaultOption()
	if conf != "" {
		configOptions.ConfigFileName = conf
	}

	switch command {
	case "version":
		fmt.Printf("ey-berks version is: %s", Version)
	case "help":
		fmt.Print(usage())
		return 0
	case "config":
		fmt.Printf("Creating a sample configuration file, %s at %s\n", configOptions.ConfigFileName, path)

		if config.IsExistConfigFile(path, configOptions) {
			fmt.Println("Error: The configration file alrady exists")
			return 1
		}

		if err := config.Create(path, configOptions); err != nil {
			fmt.Println(err)
			return 1
		}
	case "update-cache":
		fmt.Println("Updatint cookbook caches")
		berks := config.Parse(path, configOptions)
		if err := gather.Gather(path, berks); err != nil {
			fmt.Println(err)
		}
	case "create-main-recipe":
		berks := config.Parse(path, configOptions)
		list := author.CreateMainRecipe(berks)
		if err := author.CreateFile(path, list); err != nil {
			fmt.Printf("error: %v\n", err)
		}
	case "copy-recipes":
		berks := config.Parse(path, configOptions)
		if err := gather.Copy(path, berks); err != nil {
			fmt.Printf("error: %v\n", err)
		}
	case "compile":
		berks := config.Parse(path, configOptions)

		if err := gather.Gather(path, berks); err != nil {
			fmt.Println(err)
		}

		list := author.CreateMainRecipe(berks)
		if err := author.CreateFile(path, list); err != nil {
			fmt.Printf("error: %v\n", err)
		}

		if err := gather.Copy(path, berks); err != nil {
			fmt.Printf("error: %v\n", err)
		}
	default:
		fmt.Println("The command doesn't exist.Please check ey-berks help.")
		return 0
	}
	return 0
}

func parseArgs(args map[string]interface{}) (command string, path string, config string) {
	command = args["<command>"].(string)

	if v := args["<path>"]; v != nil {
		path = v.(string)
	}
	if v := args["--config"]; v != nil {
		config = v.(string)
	}
	return
}
