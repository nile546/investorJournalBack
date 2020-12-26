.PHONY: run
run:
				go run ./cmd/server/main.go

.PHONY: build
build:
				go build ./cmd/server/main.go


.DEFAULT_GOAL := run