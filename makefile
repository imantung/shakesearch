BINARY_NAME=bin/sharesearch

run:
	go build -o $(BINARY_NAME) -v .
	./$(BINARY_NAME)