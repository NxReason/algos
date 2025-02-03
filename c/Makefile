EXE=algo

all: build run

rebuild: clean build

build:
	gcc main.c array.c utils.c sort/insertion.c -o algo

run:
	$(EXE).exe

clean:
	rm $(EXE).exe
