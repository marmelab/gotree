package main

import (
	"flag"
	"github.com/marmelab/gotree"
	"github.com/nsf/termbox-go"
)

var rootPath string

func init() {
	rootPath = "."
}

func main() {
	flag.Parse()
	if userpath := flag.Arg(0); userpath != "" {
		rootPath = userpath
	}

	var screen gotree.Screen = new(gotree.Termbox)
	displayer := &gotree.Displayer{screen}

	displayer.Init()
	defer func() {
		displayer.Terminate()
	}()

	navigator := gotree.NewNavigator(displayer, rootPath)
	navigator.InitDir(rootPath)

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break loop
			case termbox.KeyArrowDown:
				navigator.ChangeSelect("down")
			case termbox.KeyArrowUp:
				navigator.ChangeSelect("up")
			case termbox.KeyArrowRight:
				navigator.EnterDir()
			case termbox.KeyArrowLeft:
				navigator.LeaveDir()
			}
		}
	}
}
