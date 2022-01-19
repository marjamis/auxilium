package engine

import (
	"fmt"

	"github.com/marjamis/auxilium/internal/pkg/blackboard"
)

// ListSteps will display the details of each step from the provided array of steps.
func ListSteps(steps []blackboard.Step) {
	for index, step := range steps {
		fmt.Printf("Step %d: %s - %+v\n", index, step.Action, step)
	}
}
