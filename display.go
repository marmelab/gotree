package gotree

import (
	"github.com/nsf/termbox-go"
)

type Displayer struct {
	Screen
}

func (d Displayer) Init() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
}

func NewDisplayer(screen Screen) *Displayer {
	return &Displayer{screen}
}

func (d Displayer) Start() {
	d.Clear(termbox.ColorDefault, termbox.ColorDefault)
	d.Flush()
}

func (d Displayer) Breadcrumb(rootPath string) {
	x := 0
	for _, r := range rootPath {
		d.SetCell(x, 0, r, termbox.ColorWhite, termbox.ColorBlue)
		x += 1
	}

	d.completeLine(x, 0, termbox.ColorWhite, termbox.ColorBlue)
}

func (d Displayer) completeLine(x int, y int, fg termbox.Attribute, bg termbox.Attribute) {
	w, _ := d.Size()
	for ; x < w; x++ {
		d.SetCell(x, y, ' ', fg, bg)
	}
}

func (d Displayer) Line(lineNumber int, file File, selected bool) {
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
		d.SetCell(x, lineNumber, r, fg, bg)

		x += 1
	}

	d.completeLine(x, lineNumber, fg, bg)
}

func (d Displayer) Stop() {
	d.Flush()
}

func (d Displayer) Terminate() {
	termbox.Close()
}
