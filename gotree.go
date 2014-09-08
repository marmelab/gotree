package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	flag.Parse()
	path := "./"
	if userpath := flag.Arg(0); userpath != "" {
		path = userpath
	}

	displayDir(path, "")
}

func displayDir(path string, previousIndent string) {
	files, _ := ioutil.ReadDir(path)
	nbElements := len(files)
	indent := "├──"
	nextIndent := " │       "

	for i, f := range files {
		if i == nbElements-1 {
			indent = "└──"
			nextIndent = "       "
		}

		if f.IsDir() {
			s := []string{path, f.Name()}
			nextPath := strings.Join(s, "/")
			fmt.Println(previousIndent, indent, f.Name())

			s[0] = previousIndent
			s[1] = nextIndent
			displayDir(nextPath, strings.Join(s, ""))
		} else {
			fmt.Println(previousIndent, indent, f.Name())
		}
	}
}

func displayIndent(previousIndent string) (nextIndent string) {
	s := []string{previousIndent, "-"}
	nextIndent = strings.Join(s, "/")
	return
}

//└──
// ├──
