build:
	@mkdir -p bin
	@go build -o bin/wargame

clean:
	@rm -f bin/wargame

run: build
	@./bin/wargame

all: build

help:
	@echo "Available targets:"
	@echo "  build  - Compiles the program to bin/wargame"
	@echo "  clean  - Removes the compiled binary"
	@echo "  run    - Builds and runs the program"
	@echo "  help   - Shows this help message"