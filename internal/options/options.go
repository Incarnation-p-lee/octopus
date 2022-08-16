package options

import (
	"flag"
)

// Option indicates the key, value and description of one option.
type Option struct {
	Key         string
	Val         string
	Description string
}

// Options indicates the collection of Option.
type Options struct {
	Options []Option
}

const (
	optionDecode       = "decode"

	invalidOptionValue = "invalid option value"
)

var supportedOptions = []Option{
	Option{
		Key: optionDecode,
		Val: invalidOptionValue,
		Description: "decode the hex/bin RISC-V instruction to asm code",
	},
}

// CreateOptions will parse the cmd parameters into Options, not supported will be ignore.
func CreateOptions() Options {
	options := Options{
		Options: []Option{},
	}

	values := make([]string, len(supportedOptions))

	for i, v := range supportedOptions {
		flag.StringVar(&values[i], v.Key, v.Val, v.Description)
	}

	flag.Parse()

	for i, v := range values {
		if v != invalidOptionValue {
			options.Options = append(options.Options, Option{
				Key: supportedOptions[i].Key,
				Val: v,
				Description: supportedOptions[i].Description,
			})
		}
	}

	return options
}
