package gotree

import (
	"io/ioutil"
	"strings"
)

var currentLine int
var currentPath string
var rootPath string
var files []File

type File struct {
	isDir bool
	name  string
}

func init() {
	currentLine = 0
}

func InitDir(path string) {
	if "" == rootPath {
		rootPath = path
	}

	displayStart()

	displayBreadcrumb(path)

	first := true
	for i, f := range fetchFiles(path) {
		displayLine(i+1, f, first)

		first = false
	}

	displayStop()

	currentPath = path
	currentLine = 0
}

func ChangeSelect(position string) {
	if "down" == position && currentLine == len(files)-1 {
		return
	}

	if "up" == position && currentLine == 0 {
		return
	}

	displayLine(currentLine+1, files[currentLine], false)

	if "up" == position {
		currentLine -= 1
	} else {
		currentLine += 1
	}

	displayLine(currentLine+1, files[currentLine], true)

	displayStop()
}

func EnterDir() {
	if !files[currentLine].isDir {
		return
	}

	s := []string{currentPath, files[currentLine].name}
	path := strings.Join(s, "/")

	InitDir(path)
}

func LeaveDir() {
	if currentPath == rootPath {
		return
	}

	splitPath := strings.Split(currentPath, "/")
	max := len(splitPath) - 1
	path := strings.Join(splitPath[0:max], "/")

	InitDir(path)
}

// use generator/parallelism/coroutine
func fetchFiles(path string) []File {
	files = make([]File, 0)
	dirFiles, _ := ioutil.ReadDir(path)

	for _, f := range dirFiles {
		if f.IsDir() && strings.HasPrefix(f.Name(), ".") {
			continue
		}

		files = append(files, File{f.IsDir(), f.Name()})
	}

	return files
}
