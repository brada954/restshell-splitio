package splitio

import (
	"errors"

	"github.com/brada954/restshell/shell"
	"github.com/splitio/go-client/v6/splitio/client"
	"github.com/splitio/go-client/v6/splitio/conf"
)

var splitApiKeyName = "SPLITIO_API_KEY"
var splitApiKey = ""
var splitActiveClient *client.SplitClient

type SplitConnectCommand struct {
	// Place getopt option value pointers here
}

func SplitConnectCommandFactory() *SplitConnectCommand {
	return &SplitConnectCommand{}
}

func (cmd *SplitConnectCommand) AddOptions(set shell.CmdSet) {
	set.SetParameters("[apikey]")
	shell.AddCommonCmdOptions(set, shell.CmdDebug, shell.CmdVerbose)
}

func (cmd *SplitConnectCommand) Execute(args []string) error {
	// Validate arguments
	splitApiKey = shell.GetGlobalString(splitApiKeyName)
	if len(args) > 0 {
		splitApiKey = args[0]
	}
	if len(splitApiKey) == 0 {
		return errors.New("no api key provided")
	}

	// Execute commands
	cfg := conf.Default()
	factory, err := client.NewSplitFactory(splitApiKey, cfg)
	if err != nil {
		return err
	}
	splitActiveClient = factory.Client()
	if err := splitActiveClient.BlockUntilReady(10); err != nil {
		splitActiveClient.Destroy()
		splitActiveClient = nil
		splitApiKey = ""
		return err
	}
	return nil
}
