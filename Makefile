include config.mk
include $(wildcard config.override.mk)

.PHONY: setup
setup:
	@go mod tidy
	@go mod vendor
	@go install github.com/air-verse/air@latest

.PHONY: run
run:
	@go run main.go

.PHONY: watch
watch:
	@air
