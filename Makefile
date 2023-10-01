
build:
	go build -o articles.service src/main.go

test:
	echo "no tests yet"

dev:
	go run src/main.go


.PHONY: build test run
