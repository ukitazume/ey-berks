package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

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

func TestNoArgs(t *testing.T) {
	command := []string{""}
	actual := getLine(commandOutput(command), 1)
	expect := "Engine Yard Cloud cookbook berkshelf"
	if actual != expect {
		t.Errorf("got %v\nwant %v", actual, expect)
	}
}

func TestHelp(t *testing.T) {
	command := []string{"help"}
	actual := getLine(commandOutput(command), 1)
	expect := "Engine Yard Cloud cookbook berkshelf"
	if actual != expect {
		t.Errorf("Expected %s, Got %s", expect, actual)
	}

}
