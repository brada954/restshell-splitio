package splitio

import "github.com/brada954/restshell/shell"

var splitCommandCategory = "SplitIO"

func init() {
	AddCommands()
}

func AddCommands() {
	shell.AddCommand("splitcon", splitCommandCategory, SplitConnectCommandFactory())
	shell.AddCommand("splitget", splitCommandCategory, SplitGetCommandFactory())
}
