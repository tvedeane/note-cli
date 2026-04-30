APP := note-cli
BIN_DIR := bin

.PHONY: build
build:
	go build -o $(BIN_DIR)/$(APP) ./cmd/$(APP)

.PHONY: test
test:
	go test ./...

.PHONY: run
run:
	go run ./cmd/$(APP)

.PHONY: clean
clean:
	rm -rf $(BIN_DIR)
