thetundra: main.go internal/story/story.go internal/logger/logger.go
	@go mod tidy
	@go build -o $@

install:
	@go install github.com/xavier2910/tundragame@latest

clean: thetundra
	@rm -f thetundra