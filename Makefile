build:
	@mkdir -p bin
	@echo "Building wargame..."
	@go build -o bin/wargame

clean:
	@echo "Cleaning up..."
	@rm -f bin/wargame

run: build
	@echo "Running wargame..."
	@./bin/wargame

all: build

help:
	@echo "Available targets:"
	@echo "  build  - Compiles the program to bin/wargame"
	@echo "  clean  - Removes the compiled binary"
	@echo "  run    - Builds and runs the program"
	@echo "  help   - Shows this help message"