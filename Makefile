.PHONY: all build

all: clean compile build

#before building pls make sure that you are logged into Docker account ( in order to avoid too mamy requests error)
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
