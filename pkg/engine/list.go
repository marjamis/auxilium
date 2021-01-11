package engine

import "fmt"

// ListSteps will display the details of each step from the provided array of steps.
func ListSteps(steps []Step) {
	for index, step := range steps {
		fmt.Printf("Step %d: %s - %+v\n", index, step.Action, step)
	}
}
