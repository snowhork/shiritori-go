// Code generated by goa v3.2.6, DO NOT EDIT.
//
// shiritori HTTP client CLI support package
//
// Command:
// $ goa gen shiritori/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	shiritoric "shiritori/gen/http/shiritori/client"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `shiritori (add|words|battle)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` shiritori add --a 6789298082735250775 --b 3453783859326640228` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
	dialer goahttp.Dialer,
	shiritoriConfigurer *shiritoric.ConnConfigurer,
) (goa.Endpoint, interface{}, error) {
	var (
		shiritoriFlags = flag.NewFlagSet("shiritori", flag.ContinueOnError)

		shiritoriAddFlags = flag.NewFlagSet("add", flag.ExitOnError)
		shiritoriAddAFlag = shiritoriAddFlags.String("a", "REQUIRED", "Left operand")
		shiritoriAddBFlag = shiritoriAddFlags.String("b", "REQUIRED", "Right operand")

		shiritoriWordsFlags    = flag.NewFlagSet("words", flag.ExitOnError)
		shiritoriWordsWordFlag = shiritoriWordsFlags.String("word", "REQUIRED", "")

		shiritoriBattleFlags        = flag.NewFlagSet("battle", flag.ExitOnError)
		shiritoriBattleBattleIDFlag = shiritoriBattleFlags.String("battle-id", "REQUIRED", "")
	)
	shiritoriFlags.Usage = shiritoriUsage
	shiritoriAddFlags.Usage = shiritoriAddUsage
	shiritoriWordsFlags.Usage = shiritoriWordsUsage
	shiritoriBattleFlags.Usage = shiritoriBattleUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "shiritori":
			svcf = shiritoriFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "shiritori":
			switch epn {
			case "add":
				epf = shiritoriAddFlags

			case "words":
				epf = shiritoriWordsFlags

			case "battle":
				epf = shiritoriBattleFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "shiritori":
			c := shiritoric.NewClient(scheme, host, doer, enc, dec, restore, dialer, shiritoriConfigurer)
			switch epn {
			case "add":
				endpoint = c.Add()
				data, err = shiritoric.BuildAddPayload(*shiritoriAddAFlag, *shiritoriAddBFlag)
			case "words":
				endpoint = c.Words()
				data, err = shiritoric.BuildWordsPayload(*shiritoriWordsWordFlag)
			case "battle":
				endpoint = c.Battle()
				data, err = shiritoric.BuildBattlePayload(*shiritoriBattleBattleIDFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// shiritoriUsage displays the usage of the shiritori command and its
// subcommands.
func shiritoriUsage() {
	fmt.Fprintf(os.Stderr, `The calc service performs operations on numbers
Usage:
    %s [globalflags] shiritori COMMAND [flags]

COMMAND:
    add: Add implements add.
    words: Words implements words.
    battle: Battle implements battle.

Additional help:
    %s shiritori COMMAND --help
`, os.Args[0], os.Args[0])
}
func shiritoriAddUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] shiritori add -a INT -b INT

Add implements add.
    -a INT: Left operand
    -b INT: Right operand

Example:
    `+os.Args[0]+` shiritori add --a 6789298082735250775 --b 3453783859326640228
`, os.Args[0])
}

func shiritoriWordsUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] shiritori words -word STRING

Words implements words.
    -word STRING: 

Example:
    `+os.Args[0]+` shiritori words --word "Beatae mollitia officia."
`, os.Args[0])
}

func shiritoriBattleUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] shiritori battle -battle-id STRING

Battle implements battle.
    -battle-id STRING: 

Example:
    `+os.Args[0]+` shiritori battle --battle-id "Illo molestiae et odio et aut animi."
`, os.Args[0])
}
