package main

import (
	"fmt"

	prompt "github.com/c-bata/go-prompt"
)

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "Canalwalk DS (0032)", Description: "Canalwalk Direct Service"},
		{Text: "Centurion DS (0005)", Description: "Centurion Direct Service"},
		{Text: "Cresta DS (0001)", Description: "Cresta Direct Service"},
		{Text: "Cresta EXP (0002)", Description: "Cresta Express Service"},
		{Text: "Eastgate EXP (0081)", Description: "Eastgate Express Service"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func main() {
	fmt.Println("Please select a branch.")
	t := prompt.Input("> ", completer)
	fmt.Println("You selected " + t)
}
