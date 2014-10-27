package gotree

import (
	"github.com/nsf/termbox-go"
)

func DisplayInit() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
}

func displayStart() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
}

func displayBreadcrumb(rootPath string) {
	x := 0
	for _, r := range rootPath {
		termbox.SetCell(x, 0, r, termbox.ColorWhite, termbox.ColorBlue)
		x += 1
	}

	completeLine(x, 0, termbox.ColorWhite, termbox.ColorBlue)
}

func completeLine(x int, y int, fg termbox.Attribute, bg termbox.Attribute) {
	w, _ := termbox.Size()
	for ; x < w; x++ {
		termbox.SetCell(x, y, ' ', fg, bg)
	}
}

func displayLine(lineNumber int, file File, selected bool) {
	bg := termbox.ColorDefault
	if selected {
		bg = termbox.ColorMagenta
	}

	fg := termbox.ColorDefault
	if file.isDir {
		fg = termbox.ColorGreen
	}

	x := 0
	for _, r := range file.name {
		termbox.SetCell(x, lineNumber, r, fg, bg)

		x += 1
	}

	completeLine(x, lineNumber, fg, bg)
}

func displayStop() {
	termbox.Flush()
}

func DisplayTerminate() {
	termbox.Close()
}
