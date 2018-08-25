markdown-cli
============
 
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/pschlump/Go-FTL/master/LICENSE)

Markdown-CLI - A simple command line interface to convert markdown to html in Go (golang)

This is a really simple program to convert from Markdown to HTML.  Options are:

	-i		--input		Input file
	-o		--output	Output file
	-p		--pre		Prefix the output with the specified string
	-P 		--post		Append the output with the specified string
	-c		--cfg		JSON config file that can set the above from a file.

Configuration can also be taken from a JSON file.  For Example, to set Pre and Post values:

	{
		"Pre": "{% body %}\n",
		"Post": "{% bodyend %}\n"
	}

Install
-------------

To Install (for Mac/Linux):
```
	$ cd ~/go/src/github.com/
	$ mkdir pschlump
	$ cd pschlump
	$ git pull https://github.com/pschlump/markdown-cli.git
	$ go get
	$ cd markdown-cli
	$ go build
```

You should have a markdown-cli executable.

Test
------------------

To test this run 

	$ make test1

License
-------

MIT - See LICENSE file.



