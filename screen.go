package gotree

import "github.com/nsf/termbox-go"

// Thanks to peco/peco

// Screen hides termbox from tne consuming code so that
// it can be swapped out for testing
type Screen interface {
	Clear(termbox.Attribute, termbox.Attribute) error
	Flush() error
	SetCell(int, int, rune, termbox.Attribute, termbox.Attribute)
	Size() (int, int)
}

// Termbox just hands out the processing to the termbox library
type Termbox struct{}

func (t Termbox) Clear(fg, bg termbox.Attribute) error {
	return termbox.Clear(fg, bg)
}

func (t Termbox) Flush() error {
	return termbox.Flush()
}

func (t Termbox) SetCell(x, y int, ch rune, fg, bg termbox.Attribute) {
	termbox.SetCell(x, y, ch, fg, bg)
}

func (t Termbox) Size() (int, int) {
	return termbox.Size()
}
