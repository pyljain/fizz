package main

import (
	"sort"
	"strings"
	"sync"

	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/pterm/pterm"
)

const maxNumMatchesPerList = 3

func updateUserInterface(lists []*cList, searchString string, area *pterm.AreaPrinter) {
	resultsChan := make(chan string)
	wg := sync.WaitGroup{}

	for _, list := range lists {
		wg.Add(1)

		go func() {
			contents := strings.Builder{}
			matches := fuzzy.RankFind(searchString, list.files)
			sort.Sort(matches)

			if len(matches) > 0 {
				r := len(matches)
				if r > maxNumMatchesPerList {
					r = maxNumMatchesPerList
				}
				for _, match := range matches[0:r] {
					contents.WriteString(match.Target + "\n")
				}
			}

			resultsChan <- contents.String()
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	output := strings.Builder{}
	output.WriteString("YOU SEARCHED FOR: " + searchString + "\n\n")
	for content := range resultsChan {
		output.WriteString(content)
		area.Update(output.String())
	}
}
