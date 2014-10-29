install:
	go get -u github.com/marmelab/gotree
	go install github.com/marmelab/gotree/cmd/gotree

test:
	go test
