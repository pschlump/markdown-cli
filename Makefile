
all:
	go build

test: test1

test1: all
	markdown-cli -i testdata/a.md -o a.html
	diff a.html testdata/a.html

test2: all
	markdown-cli -o a.html testdata/a.md
	diff a.html testdata/a.html

install: all
	cp markdown-cli ${HOME}/bin

