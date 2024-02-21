thetundra: main.go
	@go mod tidy
	@go build -o $@ main.go

install:
	@go install github.com/xavier2910/tundragame@latest

clean:
	@rm -rf build