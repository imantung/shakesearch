BINARY_NAME=bin/sharesearch

run:
	go build -o $(BINARY_NAME) -v .
	./$(BINARY_NAME)

test:
	go test ./... -v

benchmark:
	go test -benchmem -run=^$$ -bench . pulley.com/shakesearch/internal/app -v