include config.mk
include $(wildcard config.override.mk)

.PHONY: run
run:
	@go run main.go

.PHONY: watch
watch:
	@go install github.com/air-verse/air@latest
	@air
