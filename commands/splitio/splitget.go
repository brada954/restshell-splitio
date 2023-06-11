package splitio

import (
	"errors"
	"fmt"
	"strings"

	"github.com/brada954/restshell/shell"
)

type SplitGetCommand struct {
	withConfig    *bool
	attributeList *shell.StringList
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
	cmd.attributeList = set.StringListLong("attribute", 'a', "Key/value Attribute (e.g. env=Production)")

	// Add command helpers for verbose, debug, restclient and output formatting
	shell.AddCommonCmdOptions(set, shell.CmdDebug, shell.CmdVerbose)
}

func (cmd *SplitGetCommand) Execute(args []string) error {
	// Validate arguments
	if len(args) == 0 {
		return errors.New("missing split name")
	} else if len(args) > 1 {
		return errors.New("invalid number of arguemnts provided")
	}
	if len(splitApiKey) == 0 {
		return errors.New("split api key has not been set")
	}
	if splitActiveClient == nil {
		return errors.New("split client is not connected, use splitcon command")
	}

	var attributes map[string]interface{}
	attributes, err := cmd.attributesToMap()
	if err != nil {
		return err
	}

	// Execute commands
	if *cmd.withConfig == false {
		treatment := splitActiveClient.Treatment(splitApiKey, args[0], attributes)
		if err := isValidTreatment(treatment); err != nil {
			return err
		}
		return outputTreatment(treatment)
	} else {
		treatment := splitActiveClient.TreatmentWithConfig(splitApiKey, args[0], attributes)
		if err := isValidTreatment(treatment.Treatment); err != nil {
			return err
		}
		if treatment.Config == nil {
			return errors.New("treatment is without configuration")
		}
		return outputTreatmentWithConfig(treatment.Treatment, *treatment.Config)
	}
}

func (cmd *SplitGetCommand) attributesToMap() (map[string]interface{}, error) {
	result := make(map[string]interface{}, 0)
	for _, nvp := range cmd.attributeList.Values {
		pair := strings.Split(nvp, "=")
		if len(pair) != 2 {
			return nil, fmt.Errorf("unable to parse key/value attributes: %s", nvp)
		}
		key := pair[0]
		value := pair[1]
		result[key] = value
	}
	return result, nil
}

func isValidTreatment(treatment string) error {
	if strings.ToUpper(treatment) == "CONTROL" {
		return errors.New("control treatment returned")
	}
	return nil
}

func outputTreatment(treatment string) error {
	fmt.Fprintln(shell.ConsoleWriter(), treatment)
	return shell.PushText("text/plain", treatment, nil)
}

func outputTreatmentWithConfig(treatment string, config string) error {
	fmt.Fprintf(shell.ConsoleWriter(), "Treatment: %s\n%s\n", treatment, config)
	return shell.PushResult(*(shell.NewJSONResult(config)))
}
