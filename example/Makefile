example: $(wildcard **/*.go) $(wildcard *.go)
	go mod download
	go build -o example

.PHONY: run
run: example
	. ./.env && ./example

.PHONY: clean
clean:
	rm -f example
