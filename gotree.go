package main

import (
	"flag"
	"fmt"
	"github.com/nsf/termbox-go"
	"io/ioutil"
	"strings"
)

var y int

func main() {
	y = 0
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer func() {
		fmt.Println("Close Termbox'")
		termbox.Close()
	}()

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	flag.Parse()
	path := "./"
	if userpath := flag.Arg(0); userpath != "" {
		path = userpath
	}
	displayDir(path, "")
loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break loop
			}
		}
	}
}

func displayDir(path string, previousIndent string) {

	files, _ := ioutil.ReadDir(path)
	nbElements := len(files)
	indent := "├── "
	nextIndent := "│       "

	for y, f := range files {
		if y == nbElements-1 {
			indent = "└── "
			nextIndent = "       "
		}
		displayLineInTermbox(fmt.Sprintf("%s%s%s", previousIndent, indent, f.Name()))
		if f.IsDir() {
			s := []string{path, f.Name()}
			nextPath := strings.Join(s, "/")

			s[0] = previousIndent
			s[1] = nextIndent
			displayDir(nextPath, fmt.Sprintf("%s%s", previousIndent, nextIndent))
		}
	}
	termbox.Flush()
}

func displayLineInTermbox(name string) {
	x := 0
	for _, r := range name {
		termbox.SetCell(x, y, r, termbox.ColorDefault, termbox.ColorDefault)
		x += 1
	}
	y += 1
}
