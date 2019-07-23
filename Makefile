.PHONY: test
test:
	go run assets-life.go testdata/deep test/deep
	go run assets-life.go testdata/file test/file
	go run assets-life.go testdata/image test/image
	go run assets-life.go testdata/index test/index
	go run assets-life.go testdata/readdir test/readdir
	go test -v ./...
