package main

import (
	"flag"
	"github.com/nsf/termbox-go"
	"io/ioutil"
	"strings"
)

var termboxLine int
var rootPath string
var currentPath string

var currentLine int
var files []File

type File struct {
	isDir bool
	name  string
}

func init() {
	termboxLine = 1
	currentLine = 0
	rootPath = "."
}

func main() {
	flag.Parse()
	if userpath := flag.Arg(0); userpath != "" {
		rootPath = userpath
	}
	currentPath = rootPath

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer func() {
		termbox.Close()
	}()

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	//rootPathTest, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	displayFirstLine(rootPath)
	displayDir(rootPath)

	termbox.Flush()

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break loop
			case termbox.KeyArrowDown:
				changeSelect("down")
			case termbox.KeyArrowUp:
				changeSelect("up")
			case termbox.KeyArrowRight:
				enterDir()
			case termbox.KeyArrowLeft:
				leaveDir()
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

func displayDir(path string) {
	dirFiles, _ := ioutil.ReadDir(path)
	first := true
	i := 0

	for _, f := range dirFiles {
		if f.IsDir() && strings.HasPrefix(f.Name(), ".") {
			continue
		}

		if 0 < i {
			first = false
		}

		displayLineInTermbox(f.Name(), f.IsDir(), first)

		files = append(files, File{f.IsDir(), f.Name()})

		i += 1
	}
}

func displayLineInTermbox(name string, isDir bool, selected bool) {
	x := 0
	for _, r := range name {
		draw(x, termboxLine, r, isDir, selected)

		x += 1
	}
	termboxLine += 1
}

func getRelativeLevelPath(folderPath string) (levelPath int) {
	folderPath = strings.Replace(folderPath, rootPath, "", 1)
	levelPath = strings.Count(folderPath, "/")

	return
}

func draw(x int, y int, r rune, isDir bool, selected bool) {
	bg := termbox.ColorDefault
	if selected {
		bg = termbox.ColorMagenta
	}

	fg := termbox.ColorDefault
	if isDir {
		fg = termbox.ColorGreen
	}

	termbox.SetCell(x, y, r, fg, bg)
}

func changeSelect(position string) {
	if "down" == position && currentLine == len(files)-1 {
		return
	}

	if "up" == position && currentLine == 0 {
		return
	}

	x := 0
	for _, r := range files[currentLine].name {
		draw(x, currentLine+1, r, files[currentLine].isDir, false)

		x += 1
	}

	if "up" == position {
		currentLine -= 1
	} else {
		currentLine += 1
	}

	x = 0
	for _, r := range files[currentLine].name {
		draw(x, currentLine+1, r, files[currentLine].isDir, true)

		x += 1
	}

	termbox.Flush()
}

func enterDir() {
	if !files[currentLine].isDir {
		return
	}

	termboxLine = 1

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()

	s := []string{currentPath, files[currentLine].name}
	currentPath = strings.Join(s, "/")

	files = make([]File, 0)

	displayFirstLine(currentPath)
	displayDir(currentPath)

	termbox.Flush()

	currentLine = 0
}

func leaveDir() {
	if currentPath == rootPath {
		return
	}

	termboxLine = 1

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()

	splitPath := strings.Split(currentPath, "/")
	max := len(splitPath) - 1
	currentPath = strings.Join(splitPath[0:max], "/")

	files = make([]File, 0)

	displayFirstLine(currentPath)
	displayDir(currentPath)

	termbox.Flush()

	currentLine = 0
}
