
all:
	go build

test1:
	go build
	markdown-cli -i testdata/a.md -o a.html
	diff a.html testdata/a.html

