.PHONY: build generate all clean lab1 lab2 lab3 lab4

all: build run

clean:
	rm -rf build

generate:
	cmake -S . -B build/

silent-build:
	cmake --build build/

build:
	@make -s silent-build 

lab1:
	./build/lab1/lab1

lab2:
	./build/lab2/lab2

lab3:
	./build/lab3/lab3

lab4:
	./build/lab4/lab4

