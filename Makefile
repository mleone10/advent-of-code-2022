.PHONY: clean test
clean:
	rm -rf bin

test: clean
	go test ./...
