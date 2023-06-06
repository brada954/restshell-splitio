package example

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/brada954/restshell/shell"
)

var (
	PLACEHOLDER_TEXT    = "{{placeholder}}"
	EXAMPLE_URL_KEY     = "Example_Service_Url"
	DEFAULT_SERVICE_URL = "https://jsonplaceholder.typicode.com"
	DEFAULT_ENTITY      = "users"
)

type ExqueryCommand struct {
	// Place getopt option value pointers here
	optionId *int
}

func NewExqueryCommand() *ExqueryCommand {
	return &ExqueryCommand{}
}

func (cmd *ExqueryCommand) AddOptions(set shell.CmdSet) {
	set.SetParameters("[entity (default: users)]")
	set.SetUsage(func() {
		set.PrintUsage(shell.ConsoleWriter())
		cmd.extendedUsage(shell.ConsoleWriter())
	})
	// Command options
	cmd.optionId = set.IntLong("id", 0, -1, "The id of the entity to request (default: all)")

	// Add command helpers for verbose, debug, restclient and output formatting
	shell.AddCommonCmdOptions(set, shell.CmdDebug, shell.CmdVerbose, shell.CmdUrl, shell.CmdRestclient, shell.CmdFormatOutput)
}

// Execute -- Exquery command to query a well known test data site and demonstrate basic output
// Use variations of -d and -v and the --out-* options to get example output formats
func (cmd *ExqueryCommand) Execute(args []string) error {
	// Validate arguments
	var entity = DEFAULT_ENTITY
	{
		if len(args) > 0 {
			entity = args[0]
		}
	}

	var url = shell.GetGlobalStringWithFallback(EXAMPLE_URL_KEY, DEFAULT_SERVICE_URL)
	{
		url = shell.GetCmdUrlValue(url)
		if len(url) == 0 {
			return errors.New("no URL is specified or configured via env variable " + EXAMPLE_URL_KEY)
		}

		url = addSegmentToUrl(entity, url)

		// Optionally add an ID
		if *cmd.optionId > 0 {
			url = addSegmentToUrl(strconv.Itoa(*cmd.optionId), url)
		}
	}

	// Execute commands
	client := shell.NewRestClientFromOptions()
	response, err := client.DoGet(nil, url)

	// By default, a summary of the result is displayed unless verbose option is selected
	// A Custom display function can be format the output however desired
	customDisplay := func(w io.Writer, r shell.Result) error {
		var fmtResult = fmt.Sprintf("Unknown response of length: %d\n", len(r.Text))

		if root, err := r.BodyMap.GetNode("$"); err == nil {
			switch t := root.(type) {
			case map[string]interface{}:
				fmtResult = fmt.Sprintf("Object with %d properties returned\n", len(t))
			case []interface{}:
				fmtResult = fmt.Sprintf("Array of %d records returned\n", len(t))
			case interface{}:
				fmtResult = "One object found\n"
			default:
			}
		}
		fmt.Fprint(shell.ConsoleWriter(), fmtResult)
		return nil
	}

	// Demonstrate default behavior if an entity ID is selected
	if *cmd.optionId > 0 {
		customDisplay = nil
	}

	// Provide a custom display handler and return errors
	return shell.RestCompletionHandler(response, err, customDisplay)
}

func (cmd *ExqueryCommand) extendedUsage(io io.Writer) {
	fmt.Println()
	fmt.Println("This command can query the following entities:")
	fmt.Println("users")
	fmt.Println("posts")
	fmt.Println("comments")
	fmt.Println("albums")
	fmt.Println("photos")
	fmt.Println()
	fmt.Println("Try example queries:")
	fmt.Println("exquery users")
	fmt.Println("exquery --id 3 users")
	fmt.Println("exquery -v posts")
}

func addSegmentToUrl(segment string, url string) string {
	if !strings.HasSuffix(url, "/") {
		url = url + "/"
	}
	return url + segment
}
