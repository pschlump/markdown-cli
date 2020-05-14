// a CLI processor for Markdown implemented using blackfirday

// (C) Philip Schlump, 2016-2018 - MIT Licensed.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pschlump/lexie/flags"
	"github.com/russross/blackfriday"
)

var opts struct {
	Input  string `short:"i" long:"input" description:"Input" default:""`
	Output string `short:"o" long:"output" description:"Output" default:"./out.html"`
	Pre    string `short:"p" long:"pre" description:"Prefix output with" default:""`
	Post   string `short:"P" long:"post" description:"Postfix output with" default:""`
	Cfg    string `short:"c" long:"cfg" description:"Json Config File" default:"./markdown-cfg.json"`
	Debug  bool   `short:"D" long:"debug" description:"Debug flag" default:"false"`
	Help0  bool   `short:"H" long:"help" description:"Help" default:"false"`
}

var Usage = func() {
	fmt.Fprintf(os.Stderr,
		`Usage of %s:
    -i | --input FN                Input file name
    -o | --output FN               Output file name
    -b | --basic                   Default false, use "basic" markdown.  If not specified use extended.
    -c | --cfg FN                  Read a configuration file.
    -p | --pre "string"            Prefix HTML output with this string
    -P | --post "string"           Postfix HTML output with this string

If -o is not specified it defaults to stdout.

-i or a file after the options is required.

`, os.Args[0])
}

func main() {

	args, err := flags.ParseArgs(&opts, os.Args)

	if err != nil {
		Usage()
		os.Exit(1)
	}

	if len(args) == 1 || opts.Help0 {
		Usage()
		os.Exit(1)
	}

	if opts.Cfg != "" {
		if Exists(opts.Cfg) {
			s, err := ioutil.ReadFile(opts.Cfg)
			if err != nil {
				fmt.Printf("Error: Unable to read JSON config file %s, Error: %s\n", opts.Cfg, err)
				os.Exit(1)
			}
			data, err := JsonStringToData(string(s))
			if err != nil {
				fmt.Printf("Error: Unable to parse JSON config file %s, Error: %s\n", opts.Cfg, err)
				os.Exit(1)
			}
			if d, ok := data["Pre"]; ok {
				if ss, ok := d.(string); ok {
					opts.Pre = ss
				}
			}
			if d, ok := data["Post"]; ok {
				if ss, ok := d.(string); ok {
					opts.Post = ss
				}
			}
			if d, ok := data["Input"]; ok {
				if ss, ok := d.(string); ok {
					opts.Input = ss
				}
			}
			if d, ok := data["Output"]; ok {
				if ss, ok := d.(string); ok {
					opts.Output = ss
				}
			}
			if d, ok := data["Debug"]; ok {
				if bb, ok := d.(bool); ok {
					opts.Debug = bb
				}
			}
		}
	}

	var input []byte
	if opts.Input != "" {
		input, err = ioutil.ReadFile(opts.Input)
	} else if len(args) > 0 {
		input, err = ioutil.ReadFile(args[1])
	}
	if err != nil {
		fmt.Printf("Error reading %s: %s\n", opts.Input, err)
		os.Exit(1)
	}
	var output []byte
	output = blackfriday.Run(input)
	if opts.Output != "" {
		err = ioutil.WriteFile(opts.Output, []byte(opts.Pre+string(output)+opts.Post), 0644)
		if err != nil {
			fmt.Printf("Error writing %s: %s\n", opts.Output, err)
			os.Exit(1)
		}
	} else {
		fmt.Printf("%s%s%s", opts.Pre, output, opts.Post)
	}
}

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func JsonStringToData(s string) (theJSON map[string]interface{}, err error) {
	err = json.Unmarshal([]byte(s), &theJSON)
	if err != nil {
		theJSON = make(map[string]interface{})
	}
	return
}
