package engine

import (
	"bufio"
	"fmt"
	"os"

	"github.com/marjamis/auxilium/internal/pkg/blackboard"
	"github.com/marjamis/auxilium/pkg/format"
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

func actionSelector(step Step, fastforward bool) {
	switch step.Action {
	case "OutputText":
		format.Print(step.Text, step.TextColour, step.TextBackgroundColour, fastforward)
	case "RunScript":
		runScript(step.Command, step.Args, step.WorkingDirectory)
	case "Makefile":
		makefile(step.FileLocation, step.Target)
	case "Clear":
		clearScreen()
		return
	case "BashShellPrompt":
		if !fastforward {
			bashShellPrompt(step.WorkingDirectory)
		} else {
			fmt.Println("Shell skipped as in we're in fastforward mode...")
		}
	}
}

// Workflow controls execution of the possible steps from the configuration file.
func Workflow(steps []Step, blackboard *blackboard.Blackboard) {
	if blackboard.FastForwardToStep != 0 {
		for _, step := range steps[:blackboard.FastForwardToStep] {
			actionSelector(step, true)
		}
		blackboard.ContinueFromStep = blackboard.FastForwardToStep
	}
	clearScreen()

	for _, step := range steps[blackboard.ContinueFromStep:] {
		actionSelector(step, false)

		// Waits for user to hit enter before going to the next step.
		bufio.NewReader(os.Stdin).ReadString('\n')
		clearScreen()
	}
}
