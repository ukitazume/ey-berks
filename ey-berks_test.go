package main

import (
	"bytes"
	"fmt"
	"github.com/docopt/docopt-go"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

func TestParseArgs(t *testing.T) {
	cmd := []string{"help", "pathpath", "--config=Berkfile"}
	args, _ := docopt.Parse(usage(), cmd, true, "", false)
	command, path, config := parseArgs(args)

	if command != "help" {
		log.Fatal("erro parse command")
	}

	if path != "pathpath" {
		log.Fatal("erro parse path")
	}

	if config != "Berkfile" {
		log.Fatalf("erro parse opts config: %v", config)
	}

}

func commandOutput(args []string) (output string) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	Command(args)

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	w.Close()
	os.Stdout = old
	output = <-outC
	return
}

func getLine(text string, line int) string {
	lines := strings.Split(text, "\n")
	if len(lines) < line+1 {
		log.Fatalf("There is no line at line %d in %s", line, text)
	}
	return lines[line]
}

func TestVersion(t *testing.T) {
	command := []string{"version"}
	actual := commandOutput(command)
	expect := fmt.Sprintf("ey-berks version is: %s", Version)
	if actual != expect {
		t.Errorf("Expected %s, Got %s", expect, actual)
	}
}

func TestHelp(t *testing.T) {
	command := []string{"help"}
	actual := getLine(commandOutput(command), 0)
	expect := "Engine Yard Cloud cookbook tool like Berkshelf"
	if actual != expect {
		t.Errorf("Expected %s, Got %s", expect, actual)
	}
}
