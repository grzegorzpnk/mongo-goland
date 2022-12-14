.PHONY: all build

all: clean compile build

build:
	@echo "Building mongoclient app"
	sudo docker build --rm -f ./build/docker/Dockerfile -t mongoclient .

compile:
	@echo "Compiling mongoclient app"
	/bin/mkdir -p bin; \
	cd src/cmd; go build -o ../../bin/mongoclient main.go; cd ../../;

clean:
	@echo "Cleaning..."
	/bin/rm -rf bin/
