package splitio

import (
	"errors"
	"fmt"
	"strings"

	"github.com/brada954/restshell/shell"
)

type SplitGetCommand struct {
	withConfig *bool
}

func SplitGetCommandFactory() *SplitGetCommand {
	return &SplitGetCommand{}
}

func (cmd *SplitGetCommand) AddOptions(set shell.CmdSet) {
	set.SetParameters("splitname")
	set.SetUsage(func() {
		set.PrintUsage(shell.ConsoleWriter())
	})

	// Command options
	cmd.withConfig = set.BoolLong("with-config", 'w', "Aggregate treatment with configuration in the result")

	// Add command helpers for verbose, debug, restclient and output formatting
	shell.AddCommonCmdOptions(set, shell.CmdDebug, shell.CmdVerbose, shell.CmdUrl, shell.CmdRestclient, shell.CmdFormatOutput)
}

func (cmd *SplitGetCommand) Execute(args []string) error {
	// Validate arguments
	if len(splitApiKey) == 0 {
		return errors.New("split api key has not been set")
	}

	if splitActiveClient == nil {
		return errors.New("split client is not connected, use splitcon command")
	}

	// Execute commands
	if *cmd.withConfig == false {
		treatment := splitActiveClient.Treatment(splitApiKey, args[0], cmd.getSplitAttributes())
		if strings.ToUpper(treatment) == "CONTROL" {
			return errors.New("control treatment returned")
		} else {
			fmt.Fprintln(shell.ConsoleWriter(), treatment)
			return shell.PushText("text/plain", treatment, nil)
		}
	} else {
		treatment := splitActiveClient.TreatmentWithConfig(splitApiKey, args[0], cmd.getSplitAttributes())
		if strings.ToUpper(treatment.Treatment) == "CONTROL" {
			return errors.New("control treatment returned")
		} else if treatment.Config == nil {
			return errors.New("treatment is without configuration")
		} else {
			fmt.Fprintf(shell.ConsoleWriter(), "Treatment: %s\n%s\n", treatment.Treatment, *treatment.Config)
			return shell.PushResult(*(shell.NewJSONResult(*treatment.Config)))
		}
	}
}

func (cmd *SplitGetCommand) getSplitAttributes() map[string]interface{} {
	return make(map[string]interface{})
}
