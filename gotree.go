package main

import (
	"flag"
	"fmt"
	"github.com/nsf/termbox-go"
	"io/ioutil"
	//"os"
	//"path/filepath"
	"strings"
)

var termboxLine int
var folderLevelMinned int
var rootPath string

func init() {
	termboxLine = 1
	folderLevelMinned = 0
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer func() {
		termbox.Close()
	}()

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	flag.Parse()
	rootPath := "./"
	if userpath := flag.Arg(0); userpath != "" {
		rootPath = userpath
	}
	//rootPathTest, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	displayFirstLine(rootPath)
	displayDir(rootPath, "")
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

func displayFirstLine(rootPath string) {
	x := 0

	for _, r := range rootPath {
		termbox.SetCell(x, 0, r, termbox.ColorWhite, termbox.ColorBlue)
		x += 1
	}
	w, _ := termbox.Size()
	for x = x; x < w; x++ {
		termbox.SetCell(x, 0, ' ', termbox.ColorWhite, termbox.ColorBlue)
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
		displayLineInTermbox(fmt.Sprintf("%s%s%s", previousIndent, indent, f.Name()), f.IsDir())
		if f.IsDir() && folderLevelMinned > 1 {
			s := []string{path, f.Name()}
			nextPath := strings.Join(s, "/")
			s[0] = previousIndent
			s[1] = nextIndent
			displayDir(nextPath, fmt.Sprintf("%s%s", previousIndent, nextIndent))
		}
	}
	termbox.Flush()
}

func displayLineInTermbox(name string, isDir bool) {
	// TODO transform name string in runes ?
	x := 0
	for _, r := range name {
		if isDir {
			termbox.SetCell(x, termboxLine, r, termbox.ColorGreen, termbox.ColorDefault)
		} else {
			termbox.SetCell(x, termboxLine, r, termbox.ColorDefault, termbox.ColorDefault)
		}

		x += 1
	}
	termboxLine += 1
}

func getRelativeLevelPath(folderPath string) (levelPath int) {
	folderPath = strings.Replace(folderPath, rootPath, "", 1)
	levelPath = strings.Count(folderPath, "/")

	return
}
