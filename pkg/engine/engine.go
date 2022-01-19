package engine

import (
	"bufio"
	"fmt"
	"os"

	"github.com/marjamis/auxilium/internal/pkg/blackboard"
	"github.com/marjamis/auxilium/pkg/format"
)

func actionSelector(step blackboard.Step, fastforward bool) {
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
func Workflow(steps []blackboard.Step, blackboard *blackboard.Blackboard) {
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
