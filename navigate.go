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

func InitDir(path string) {
	if "" == rootPath {
		rootPath = path
	}

	displayStart()

	displayBreadcrumb(path)

	c := make(chan File)
	go fetchFiles(path, c)

	first := true
	i := 0
	for f := range c {
		displayLine(i+1, f, first)

		first = false
		i += 1
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

func fetchFiles(path string, fs chan File) {
	files = make([]File, 0)
	var file File
	dirFiles, _ := ioutil.ReadDir(path)

	for _, f := range dirFiles {
		if f.IsDir() && strings.HasPrefix(f.Name(), ".") {
			continue
		}
		file = File{f.IsDir(), f.Name()}
		fs <- file

		files = append(files, file)
		// do something
	}

	close(fs)
}
