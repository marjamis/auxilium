package blackboard

import (
	"fmt"
)

// Step is the struct for all the possible information that a Step may have.
// Note: in the future perhaps I should break these out into their own structures rather than being this flat? Especially as not all combinations make sense.
type Step struct {
	Action               string   `yaml:"Action"`
	Text                 string   `yaml:"Text"`
	TextColour           string   `yaml:"TextColour,omitempty"`
	TextBackgroundColour string   `yaml:"BackgroundColour,omitempty"`
	Command              string   `yaml:"Command"`
	Args                 []string `yaml:"Args"`
	FileLocation         string   `yaml:"FileLocation"`
	Target               string   `yaml:"Target"`
	WorkingDirectory     string   `yaml:"WorkingDirectory"`
}

// Defaults contains data of default values used for details in the cluster via the blackboard.
type Defaults struct {
	BackgroundColour string `yaml:"BackgroundColour"`
	TextColour       string `yaml:"TextColour"`
	WorkingDirectory string `yaml:"WorkingDirectory"`
}

// Blackboard contains all state information used for the application and can be easily referenced as required.
type Blackboard struct {
	TutorialName      string
	ContinueFromStep  int
	FastForwardToStep int
	Defaults          Defaults
	Steps             []Step `yaml:"Steps"`
}

// BlackboardData is the default blackboard used for this package but custom versions of this Blackboard can be created.
var BlackboardData Blackboard

// PrintOutBlackboardData will print the entire struct data for the provided Blackboard object, normally the default listed above (var BlackboardData Blackboard).
func (bbd *Blackboard) PrintOutBlackboardData() {
	fmt.Printf("%+v\n", bbd)
}
