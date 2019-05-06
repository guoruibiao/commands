package commands

import (
	"bytes"
	"log"
	"os/exec"
)

/** 类似于python的commands库
 * - Run(command string, args...string) (void)
 * - GetOutput(command string, args...string) (output string)
 * - GetStatusOutput(command string, args...string) (status bool, output string)
 **/

// Commands define the structure
type Commands struct {
	Status bool
	Output string
}

// New init instance with properties stands for the status and the output
func New() *Commands {
	var command Commands
	command = Commands{
		Status: false,
		Output: "",
	}
	return &command
}

// Run commands without any output
func (c *Commands) Run(command string, args ...string) {
	cmd := exec.Command(command, args...)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

// GetOutput run commands without status, but the output
func (c *Commands) GetOutput(command string, args ...string) (output string) {
	bytes, err := exec.Command(command, args...).Output()
	if err != nil {
		c.Output = err.Error()
	} else {
		c.Output = string(bytes)
	}
	return c.Output
}

// GetStatusOutput run commands with status and output
func (c *Commands) GetStatusOutput(command string, args ...string) (status bool, output string) {
	cmd := exec.Command(command, args...)
	var bufferout, bufferin, buffererr bytes.Buffer
	cmd.Stdout = &bufferout
	cmd.Stdin = &bufferin
	cmd.Stderr = &buffererr

	// do the commands
	err := cmd.Run()
	c.Status = cmd.ProcessState.Success()
	if err != nil {
		c.Output = err.Error()
	} else {
		c.Output = bufferout.String()
	}
	return c.Status, c.Output
}
