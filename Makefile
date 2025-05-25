test:
	go test -v -count 1 ./... 
run:
	go run ./cmd/shortener

.PHONY: test run