package splitio

import (
	"github.com/brada954/restshell/shell"
)

type SplitCloseCommand struct {
	// Place getopt option value pointers here
}

func SplitCloseCommandFactory() *SplitCloseCommand {
	return &SplitCloseCommand{}
}

func (cmd *SplitCloseCommand) AddOptions(set shell.CmdSet) {
	shell.AddCommonCmdOptions(set, shell.CmdDebug, shell.CmdVerbose)
}

func (cmd *SplitCloseCommand) Execute(args []string) error {
	// Validate arguments

	// Execute commands
	if splitActiveClient != nil {
		splitActiveClient.Destroy()
		splitActiveClient = nil
		splitApiKey = ""
	}
	return nil
}
