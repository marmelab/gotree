package gotree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitDir(t *testing.T) {
	assert := assert.New(t)

	// assert equality
	assert.Equal(currentLine, 0, "CurrentLIne is 0 on instantiation")
	currentLine = 3
	InitDir("fake/path")
	assert.Equal(currentLine, 0, "CurrentLIne is 0 after InitDir()")
	assert.Equal(currentPath, "fake/path", "CurrentPath is set to new value by InitDir()")
}

// func TestChangeSelect(t *testing.T) {
// 	assert := assert.New(t)
// 	currentLine = 3
// 	files = make([]File, 0)
// 	files = append(files, File{false, "fileOne"})
// 	files = append(files, File{false, "fileTwo"})
// 	files = append(files, File{true, "folder"})
// 	ChangeSelect("up")
// 	assert.Equal(currentLine, 1, "CurrentLine is increment by one with ChangeSelect up")
// 	ChangeSelect("up")
// 	assert.Equal(currentLine, 2, "CurrentLine is increment by one with ChangeSelect up, even on a folder")
// 	ChangeSelect("up")
// 	assert.Equal(currentLine, 2, "CurrentLine don't increment by one the last folder entry")
// 	ChangeSelect("down")
// 	assert.Equal(currentLine, 1, "CurrentLine is decrement by one with ChangeSelect down")
// }
