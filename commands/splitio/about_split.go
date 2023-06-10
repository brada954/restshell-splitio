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
	Title:       "Split I/O",
	Description: "Split I/O commands for access Split I/O Feature flags and configuration",
	About: `Split I/O is a SAAS service with an API for accessing configuration. Applications
make API calls to get the value of configuration settings during application runtime.
The Split I/O commands allow restshell to qury the API to get the configuration directly
as an application.

Use the SPLITCON command to make a connection to split. The command takes an API key as a 
parameter:

	SPLITCON my_api_key

The SPLITGET command can query the value of the Treatment (i.e. on, off, etc) or with options
query the configuration. 

    SPLITGET --with-config enableMyWidgetFeature

When the split configuration is JSON it will be stored as a result and can be interogated
like any other JSON REST response.

The SPLITDISC command disconnects the Split I/O client. Once the split is connected, the 
client connection stays open as an application would use Split I/O
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
