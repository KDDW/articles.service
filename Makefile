NAME := articles


build:
	go build -o $(NAME) src/main.go

test:
	echo "no tests yet"

dev:
	go run src/main.go

clean:
	rm -f $(NAME)


.PHONY: build test run
