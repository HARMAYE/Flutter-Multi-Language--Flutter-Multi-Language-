all: build install 

build:
	go get -t
	go build

install:
	go install

test: build install
