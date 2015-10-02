package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/ukitazume/ey-berks/attribute"
	"github.com/ukitazume/ey-berks/author"
	"github.com/ukitazume/ey-berks/config"
	"github.com/ukitazume/ey-berks/gather"
)

const (
	Version = "0.0.1"
)

func main() {
	os.Exit(Command(os.Args[1:]))
}

func usage() string {
	return `Engine Yard Cloud cookbook tool like Berkshelf

Usage:
  ey-berks config <path>                                               : make a sample configuration file (default path=$PWD, --config=EyBerks)
  ey-berks compile <path> --config=<path to EyBerks>                   : update cahce,  write a main/recipes and gather recipe to the cookbooks directory
  ey-berks update-cache                                                : update cache of remote repositories cookbooks
  ey-berks create-main-recipe <path>                                   : create main recipes from the configration file
  ey-berks copy-recipes <path>                                         : copy recipes from the cache dir to the cookbooks/ directory
  ey-berks clear <path>                                                : remove EyBerksfile and cookbooks directory
  ey-berks gather-attr <path> --from=</path/cookbooks>         : gather attbiutes from cookbook directory
  ey-berks apply-attr <path> ---from=<attributes directory> : apply attbiutes for cookbook directory
  ey-berks help                                                        : show this help
  ey-berks version                                                     : show the version
`
}

func Command(argv []string) int {
	command, args, options := parseArgs(argv)

	configOptions := config.DefaultOption()
	if options["config"] != "" {
		configOptions.ConfigFileName = options["config"]
	}
	path := "./"
	if len(args) > 0 {
		path = args[0]
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
			removeDirs(
				filepath.Join(path, configOptions.ConfigFileName),
				filepath.Join(path, configOptions.TargetDirName),
			)
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
	case "apply-attr":
		if _, ok := options["from"]; !ok {
			fmt.Println("require --from=directory path")
			return 1
		}
		attrOptions := attribute.DefaultOptions()
		if err := attribute.Apply(path, options["from"], attrOptions); err != nil {
			fmt.Printf("error: %v\n", err)
		}
	case "gather-attr":
		if _, ok := options["from"]; !ok {
			fmt.Println("require --from=directory path")
			return 1
		}
		attrOptions := attribute.DefaultOptions()
		if err := attribute.Gather(options["from"], path, attrOptions); err != nil {
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

func removeDirs(pathes ...string) {
	for _, path := range pathes {
		if err := os.RemoveAll(path); err != nil {
			fmt.Printf("error: dont' remove because %v\n", err)
		}
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

func parseArgs(argv []string) (command string, args []string, options map[string]string) {
	options = map[string]string{}
	if len(argv) >= 1 {
		command = argv[0]
	} else {
		command = "help"
		return
	}
	reg, _ := regexp.Compile("--([a-z]+)=(.+)")
	for _, value := range argv[1:] {
		if match := reg.FindStringSubmatch(value); len(match) == 3 {
			key := match[1]
			options[key] = match[2]
		} else {
			args = append(args, value)
		}
	}
	return
}
