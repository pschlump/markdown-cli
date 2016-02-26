
all:
	go build

test1: all
	markdown-cli -i testdata/a.md -o a.html
	diff a.html testdata/a.html

install: all
	echo cp markdown.cli ${HOME}/bin

