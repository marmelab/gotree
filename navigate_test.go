package gotree

import (
	"github.com/nsf/termbox-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type DummyScreen struct {
	mock.Mock
}

func (ds *DummyScreen) Clear(fg, bg termbox.Attribute) error {
	args := ds.Called(fg, bg)

	return args.Error(0)
}

func (ds *DummyScreen) Flush() error {
	args := ds.Called()

	return args.Error(0)
}

func (ds *DummyScreen) SetCell(x, y int, ch rune, fg, bg termbox.Attribute) {
	ds.Called(x, y, ch, fg, bg)
}

func (ds *DummyScreen) Size() (int, int) {
	args := ds.Called()

	return args.Int(0), args.Int(1)
}

func TestInitDir(t *testing.T) {
	assert := assert.New(t)

	screen := new(DummyScreen)
	displayer := NewDisplayer(screen)
	navigator := NewNavigator(displayer, "fake")

	screen.On("Clear", 0, 0).Return(nil)
	screen.On("Flush").Return(nil)
	screen.On("SetCell", 0, 0, 102, 8, 5).Return()
	screen.On("SetCell", 1, 0, 97, 8, 5).Return()
	screen.On("SetCell", 2, 0, 107, 8, 5).Return()
	screen.On("SetCell", 3, 0, 101, 8, 5).Return()
	screen.On("SetCell", 4, 0, 47, 8, 5).Return()
	screen.On("SetCell", 5, 0, 100, 8, 5).Return()
	screen.On("SetCell", 6, 0, 105, 8, 5).Return()
	screen.On("SetCell", 7, 0, 114, 8, 5).Return()
	screen.On("Size").Return(1, 1)

	// assert equality
	assert.Equal(navigator.currentLine, 0, "currentLne is 0 on instantiation")
	assert.Equal(navigator.rootPath, "fake", "rootPath is set to new value by InitDir()")
	assert.Equal(navigator.currentPath, "fake", "currentPath is set to new value by InitDir()")
	navigator.currentLine = 3
	navigator.InitDir("fake/dir")
	assert.Equal(navigator.currentLine, 0, "currentLine is 0 after InitDir()")
	assert.Equal(navigator.currentPath, "fake/dir", "currentPath is set to new value by InitDir()")
}

func TestChangeSelect(t *testing.T) {
	assert := assert.New(t)

	screen := new(DummyScreen)

	screen.On("SetCell", 0, 3, 102, 3, 0).Return()
	screen.On("SetCell", 1, 3, 111, 3, 0).Return()
	screen.On("SetCell", 2, 3, 108, 3, 0).Return()
	screen.On("SetCell", 3, 3, 100, 3, 0).Return()
	screen.On("SetCell", 4, 3, 101, 3, 0).Return()
	screen.On("SetCell", 5, 3, 114, 3, 0).Return()
	screen.On("Size").Return(1, 1)
	screen.On("SetCell", 0, 2, 102, 0, 6).Return()
	screen.On("SetCell", 1, 2, 105, 0, 6).Return()
	screen.On("SetCell", 2, 2, 108, 0, 6).Return()
	screen.On("SetCell", 3, 2, 101, 0, 6).Return()
	screen.On("SetCell", 4, 2, 50, 0, 6).Return()
	screen.On("Size").Return(1, 1)
	screen.On("Flush").Return(nil)

	displayer := NewDisplayer(screen)
	navigator := NewNavigator(displayer, "fake")

	navigator.currentLine = 2
	navigator.files = make([]File, 0)
	navigator.files = append(navigator.files, File{false, "file1"})
	navigator.files = append(navigator.files, File{false, "file2"})
	navigator.files = append(navigator.files, File{true, "folder"})

	navigator.ChangeSelect("up")
	assert.Equal(navigator.currentLine, 1, "CurrentLine is increment by one with ChangeSelect up")

	screen.AssertExpectations(t)

	// navigator.ChangeSelect("up")
	// assert.Equal(navigator.currentLine, 2, "CurrentLine is increment by one with ChangeSelect up, even on a folder")
	// navigator.ChangeSelect("up")
	// assert.Equal(navigator.currentLine, 2, "CurrentLine don't increment by one the last folder entry")
	// navigator.ChangeSelect("down")
	// assert.Equal(navigator.currentLine, 1, "CurrentLine is decrement by one with ChangeSelect down")
}
