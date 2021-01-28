BINARY_NAME=bin/sharesearch

run:
	go build -o $(BINARY_NAME) -v .
	./$(BINARY_NAME)

benchmark:
	go test -run=^$$ -bench . pulley.com/shakesearch/internal/app -v