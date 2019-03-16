GOBUILD_OPT=-mod=vendor -v
BINARY=librc
BINARY_WINDOWS=librc.exe

build:
	go build $(GOBUILD_OPT) -o $(BINARY)
build-windows:
	go build $(GOBUILD_OPT) -o $(BINARY_WINDOWS)
