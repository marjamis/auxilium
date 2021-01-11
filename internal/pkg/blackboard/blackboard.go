package blackboard

import "fmt"

// TODO am I just replicating the configuration file here? especially as there isn't any validation? If we have validation on top then Im happy. This plus the file and how these are set/validated from the configuration file and the blackboard all needs a bit of a think and reworking.
// Defaults contains data of default values used for details in the cluster via the blackboard.
type Defaults struct {
	BackgroundColour string
	TextColour       string
	WorkingDirectory string
}

// Blackboard contains all state information used for the application and can be easily referenced as required.
type Blackboard struct {
	TutorialName      string
	ContinueFromStep  int
	FastForwardToStep int
	Defaults          Defaults
}

// BlackboardData is the default blackboard used for this package but custom versions of this Blackboard can be created.
var BlackboardData Blackboard

// PrintOutBlackboardData will print the entire struct data for the provided Blackboard object, normally the default listed above (var BlackboardData Blackboard).
func (bbd *Blackboard) PrintOutBlackboardData() {
	fmt.Printf("%+v\n", bbd)
}
