current_dir := $(shell pwd)
source_root := $(current_dir)/cmd
build_root := $(current_dir)/dist

clean:
	@echo "Cleaning build directory..."
	@rm -rf $(build_root)

build: clean
	@echo "Building Luminary Base Server..."
	@go build -o $(build_root)/luminary-base-server $(source_root)/luminary-base-server.go
