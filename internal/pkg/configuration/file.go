package configuration

import (
	"github.com/marjamis/auxilium/pkg/engine"
)

// TODO better jiving of this with the blackboard as they're different but overlapping
// Defaults set in the configuration file for steps to use.
type Defaults struct {
	BackgroundColour string `yaml:"BackgroundColour"`
	TextColour       string `yaml:"TextColour"`
	WorkingDirectory string `yaml:"WorkingDirectory"`
}

// ConfigurationFile is the struct for the configuration file.
type File struct {
	TutorialName string        `yaml:"TutorialName"`
	Defaults     Defaults      `yaml:"Defaults"`
	Steps        []engine.Step `yaml:"Steps"`
}
