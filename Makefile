.PHONY: test
test:
	cd testdata && go run generatebench.go
	go run assets-life.go testdata/bench test/bench
	go run assets-life.go testdata/deep test/deep
	go run assets-life.go testdata/file test/file
	go run assets-life.go testdata/image test/image
	go run assets-life.go testdata/index test/index
	go run assets-life.go testdata/readdir test/readdir
	go test -v -bench . -benchmem ./...
	go generate ./...
