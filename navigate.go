package gotree

import (
	"io/ioutil"
	"strings"
)

type File struct {
	isDir bool
	name  string
}

type Navigator struct {
	display     *Displayer
	currentPath string
	rootPath    string
	files       []File
	currentLine int
}

func NewNavigator(display *Displayer, path string) *Navigator {
	files := make([]File, 0)

	return &Navigator{display, path, path, files, 0}
}

func (n *Navigator) InitDir(path string) {
	n.display.Start()

	n.display.Breadcrumb(path)

	c := make(chan File)
	go n.fetchFiles(path, c)

	first := true
	i := 0
	for f := range c {
		n.display.Line(i+1, f, first)

		first = false
		i += 1
	}

	n.display.Stop()

	n.currentPath = path
	n.currentLine = 0
}

func (n *Navigator) ChangeSelect(position string) {
	if "down" == position && n.currentLine == len(n.files)-1 {
		return
	}

	if "up" == position && n.currentLine == 0 {
		return
	}

	n.display.Line(n.currentLine+1, n.files[n.currentLine], false)

	if "up" == position {
		n.currentLine -= 1
	} else {
		n.currentLine += 1
	}

	n.display.Line(n.currentLine+1, n.files[n.currentLine], true)

	n.display.Stop()
}

func (n *Navigator) EnterDir() {
	if !n.files[n.currentLine].isDir {
		return
	}

	s := []string{n.currentPath, n.files[n.currentLine].name}
	path := strings.Join(s, "/")

	n.InitDir(path)
}

func (n *Navigator) LeaveDir() {
	if n.currentPath == n.rootPath {
		return
	}

	splitPath := strings.Split(n.currentPath, "/")
	max := len(splitPath) - 1
	path := strings.Join(splitPath[0:max], "/")

	n.InitDir(path)
}

func (n *Navigator) fetchFiles(path string, fs chan File) {
	n.files = make([]File, 0)
	var file File
	dirFiles, _ := ioutil.ReadDir(path)

	for _, f := range dirFiles {
		if f.IsDir() && strings.HasPrefix(f.Name(), ".") {
			continue
		}
		file = File{f.IsDir(), f.Name()}
		fs <- file

		n.files = append(n.files, file)
		// do something more complex
	}

	close(fs)
}
