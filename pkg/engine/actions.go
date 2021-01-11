package engine

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/marjamis/auxilium/internal/pkg/blackboard"
)

var (
	bbd = &blackboard.BlackboardData
)

// Clears the terminal screen. Using ANSI characters to do this rather than the 'clear' command as it's shorter in lenght and doesn't generate weird characters on the clear when you scroll up. Taken from: https://stackoverflow.com/a/22892171 :)
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

// TODO should these be in a mapping to show they're tied together?
func runScript(command string, arguments []string, workingdirectory string) {
	path, _ := exec.LookPath(command)
	args := make([]string, 1+len(arguments))

	args[0] = path
	for i, arg := range arguments {
		args[i+1] = arg
	}

	// TODO find a way to wrap this consistently as currently I have to do this over and over again
	if workingdirectory == "" && bbd.Defaults.WorkingDirectory != "" {
		workingdirectory = bbd.Defaults.WorkingDirectory
	}
	cmd := exec.Cmd{
		Path:   path,
		Args:   args,
		Dir:    workingdirectory,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	if err := cmd.Run(); err != nil {
		fmt.Printf("There was an error of: %s", err)
	}
}

func makefile(fileLocation string, target string) {
	path, err := exec.LookPath("make")
	if err != nil {
		fmt.Printf("func makefile() had an error of: ''%s'", err)
		return
	}

	cmd := exec.Cmd{
		Path: path,
		Args: []string{
			path,
			"-f",
			fileLocation,
			target,
		},
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	if err := cmd.Run(); err != nil {
		fmt.Printf("There was an error of: %s", err)
	}
}

func bashShellPrompt(workingdirectory string) {
	if workingdirectory == "" && bbd.Defaults.WorkingDirectory != "" {
		workingdirectory = bbd.Defaults.WorkingDirectory
	}
	proc, err := os.StartProcess("/bin/bash", []string{"--login"}, &os.ProcAttr{
		// Transfer stdin, stdout, and stderr to the new process
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Dir:   workingdirectory,
		// Set an environment variable to indicate the shell is in the auxilium run.
		Env: []string{
			"auxilium=true",
		},
	})
	if err != nil {
		panic(err)
	}

	// Wait until user exits the shell
	state, err := proc.Wait()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Exited shell with exit code: %s. Press enter to continue...\n", state.String())
}
