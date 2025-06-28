package main

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/charlievieth/fastwalk"
	"github.com/pterm/pterm"
)

func main() {
	dir := os.Args[1]

	if dir == "" {
		dir = "."
	}

	numCPUs := runtime.NumCPU()
	var lists = make([]*cList, numCPUs)
	for i := range lists {
		lists[i] = newcList()
	}

	currentListIndex := 0
	err := fastwalk.Walk(nil, dir, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		p, _ := filepath.Rel(dir, path)

		lists[currentListIndex].add(p)
		currentListIndex += 1
		if currentListIndex >= numCPUs-1 {
			currentListIndex = 0
		}

		return nil
	})
	if err != nil {
		log.Printf("Error occured %s", err)
		os.Exit(-1)
	}

	searchString := strings.Builder{}
	area, _ := pterm.DefaultArea.Start()
	defer area.Stop()

	area.Update("Search: ")

	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		switch key.Code {
		case keys.CtrlC, keys.Escape:
			return true, nil
		case keys.Backspace, keys.Delete:
			var truncatedSearchString string
			originalSearchString := searchString.String()
			if len(originalSearchString) > 0 {
				truncatedSearchString = originalSearchString[0 : len(originalSearchString)-1]
			}
			searchString.Reset()
			searchString.WriteString(truncatedSearchString)
			updateUserInterface(lists, searchString.String(), area)
		default:
			searchString.WriteString(string(key.String()))
			updateUserInterface(lists, searchString.String(), area)
		}

		return false, nil // Return false to continue listening
	})

}
