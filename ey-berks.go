package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"github.com/ukitazume/ey-berks/author"
	"github.com/ukitazume/ey-berks/config"
	"github.com/ukitazume/ey-berks/gather"
	"log"
	"os"
	"path/filepath"
)

const (
	Version = "0.0.1"
)

func main() {
	os.Exit(Command(os.Args[1:]))
}

func usage() string {
	return `Engine Yard Cloud cookbook tool like Berkshelf

Usage: ey-berks <command> [<path>] [--config=<config>]

ey-berks config <path> [--config=<config>]             : make a sample configuration file
ey-berks compile <path> [--config=<config>]            : update cahce,  write a main/recipes and gather recipe to the cookbooks directory
ey-berks update-cache [--config=<config>]              : update cache of remote repositories cookbooks
ey-berks create-main-recipe <path> [--config=<config>] : create main recipes from the configration file
ey-berks copy-recipes <path> [--config=<config>]       : copy recipes from the cache dir to the cookbooks/ directory
ey-berks clear <path>                                  : remove EyBerksfile and cookbooks directory
ey-berks help                                          : show this help
ey-berks version                                       : show the version
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
	case "clear":
		fmt.Println("remove cookbooks and %s at %s ? [y, yes|n, no]", configOptions.ConfigFileName, path)
		if askForConfirmation() {
			fmt.Println("removing cookbooks/ and %s", configOptions.ConfigFileName)
			cookbookPath := filepath.Join(path, configOptions.ConfigFileName)
			configPath := filepath.Join(path, configOptions.TargetDirName)
			if err := os.RemoveAll(cookbookPath); err != nil {
				fmt.Printf("error: dont' remove because %v\n", err)
			}
			if err := os.RemoveAll(configPath); err != nil {
				fmt.Printf("error: dont' remove because %v\n", err)
			}
			fmt.Println("removed")
		}
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

func askForConfirmation() bool {
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}
	okayResponses := []string{"y", "Y", "yes", "Yes", "YES"}
	nokayResponses := []string{"n", "N", "no", "No", "NO"}
	if containsString(okayResponses, response) {
		return true
	} else if containsString(nokayResponses, response) {
		return false
	} else {
		fmt.Println("Please type yes or no and then press enter:")
		return askForConfirmation()
	}
}

func containsString(slice []string, element string) bool {
	return !(posString(slice, element) == -1)
}

func posString(slice []string, element string) int {
	for index, elem := range slice {
		if elem == element {
			return index
		}
	}
	return -1
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
