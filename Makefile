all: build run

build:
	go build -o get-wx

run:
	./get-wx

install:
	cp get-wx ~/bin/get-wx