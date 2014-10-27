package main

import (
	"flag"
	"github.com/nsf/termbox-go"
	"gotree"
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

	gotree.DisplayInit()
	defer func() {
		gotree.DisplayTerminate()
	}()

	gotree.InitDir(rootPath)

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break loop
			case termbox.KeyArrowDown:
				gotree.ChangeSelect("down")
			case termbox.KeyArrowUp:
				gotree.ChangeSelect("up")
			case termbox.KeyArrowRight:
				gotree.EnterDir()
			case termbox.KeyArrowLeft:
				gotree.LeaveDir()
			}
		}
	}
}
