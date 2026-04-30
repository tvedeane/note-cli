APP := note
BIN_DIR := bin
RUN_ARGS := $(filter-out run,$(MAKECMDGOALS))

.PHONY: build
build:
	go build -o $(BIN_DIR)/$(APP) ./cmd/$(APP)

.PHONY: test
test:
	go test ./...

.PHONY: run
run:
	go run ./cmd/$(APP) $(RUN_ARGS)

.PHONY: clean
clean:
	rm -rf $(BIN_DIR)

%:
	@:
