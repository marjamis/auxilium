package format

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/logrusorgru/aurora"

	"github.com/marjamis/auxilium/internal/pkg/blackboard"
)

const (
	variationInTypingSpeed int = 15

	slowTypingSpeed    int = 100
	averageTypingSpeed int = 50
	fastTypingSpeed    int = 10
)

var (
	bbd = &blackboard.BlackboardData
)

//Print will take in a string, text color, background text color and if slow writing (simulating typing) is disabled and with this information generate text strings.
func Print(text string, textColour string, textBackgroundColor string, disableSlowWrites bool) {
	colour := colour(textColour, textBackgroundColor)
	for _, char := range text {
		fmt.Print(aurora.Colorize(string(char), colour))

		delay := rand.Intn(variationInTypingSpeed) + fastTypingSpeed
		if disableSlowWrites {
			delay = 0
		}
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
}

func colour(textColour string, textBackgroundColour string) aurora.Color {
	var c aurora.Color

	// Read the defaults from the Blackboard and use those if not explicitly set in the configuration file.
	if textColour == "" && bbd.Defaults.TextColour != "" {
		textColour = bbd.Defaults.TextColour
	}
	if textBackgroundColour == "" && bbd.Defaults.BackgroundColour != "" {
		textBackgroundColour = bbd.Defaults.BackgroundColour
	}

	switch strings.ToLower(textColour) {
	case "blue":
		c |= aurora.BlueFg
	case "green":
		c |= aurora.GreenFg
	case "red":
		c |= aurora.RedFg
	case "yellow":
		c |= aurora.YellowFg
	case "black":
		c |= aurora.BlackFg
	case "white":
		c |= aurora.WhiteFg
	}
	switch strings.ToLower(textBackgroundColour) {
	case "blue":
		c |= aurora.BlueBg
	case "green":
		c |= aurora.GreenBg
	case "red":
		c |= aurora.RedBg
	case "yellow":
		c |= aurora.YellowBg
	case "black":
		c |= aurora.BlackBg
	case "white":
		c |= aurora.WhiteBg
	}

	return c
}
