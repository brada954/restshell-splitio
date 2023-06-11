package splitio

import (
	"fmt"
	"io"
)

type SplitIOTopic struct {
	Key         string
	Title       string
	Description string
	About       string
}

var splitTopic = &SplitIOTopic{
	Key:         "SPLIT",
	Title:       "Split.io",
	Description: "Split commands for access Split.io Feature flags and configuration",
	About: `Split.io is a SAAS service with an API for accessing configuration. Applications
make API calls to get the value of configuration settings during application runtime.
The Split commands allow restshell to query the API to get the configuration directly
as an application.

Connect to Split using the SPLITCON command maintains the connection until closed.

	SPLITCON my_api_key

SPLITCON will get a default API key from the configuration variable 'SPLITIO_API_KEY' if not
specified with the command.

Get Split I/O treatments or configurations with the SPLITGET command:

    SPLITGET --with-config my_split_name

Attributes can be provided as paramters for targeting rules. The split configuration is
stored as a result and can be interogated like any other JSON response.

Close the Split.io connection with the SPLITCLOSE command:

	SPLITCLOSE
`,
}

// NewSubstitutionTopic -- return a topic structure for help about Substitutions
func NewSplitTopic() *SplitIOTopic {
	return splitTopic
}

// GetKey -- return the key to the substitution help topic
func (a *SplitIOTopic) GetKey() string {
	return a.Key
}

// GetTitle() -- return the title of the substitution about topic
func (a *SplitIOTopic) GetTitle() string {
	return a.Title
}

// GetDescription -- return the description of substitution about topic
func (a *SplitIOTopic) GetDescription() string {
	return a.Description
}

// WriteAbout -- write the substitution about topic to the provided writer
func (a *SplitIOTopic) WriteAbout(o io.Writer) error {
	fmt.Fprintf(o, a.About)
	return nil
}

// WriteSubTopic -- Write the Subtopic about information
func (a *SplitIOTopic) WriteSubTopic(o io.Writer, fname string) error {
	return nil
}
