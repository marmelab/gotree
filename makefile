install:
	- go get -u github.com/nsf/termbox-go
	- go get -u github.com/stretchr/testify

test:
	- go test

installInCli:
	- go get -u github.com/marmelab/gotree
	- go install github.com/marmelab/gotree/cmd/gotree
