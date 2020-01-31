package commands

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
	)

/** 类似于python的commands库
 * - Run(command string, args...string) (void)
 * - GetOutput(command string, args...string) (output string)
 * - GetStatusOutput(command string, args...string) (status bool, output string)
 **/

 // 执行连续的shell命令时, 需要注意指定执行路径和参数, 否则运行出错
 // 比如 /bin/sh -c multi-commands-string

// Commands define the structure
type Commands struct {
	Status bool
	Output string
	Error error
}

// New init instance with properties stands for the status and the output
func New() *Commands {
	var command Commands
	command = Commands{
		Status: false,
		Output: "",
		Error: nil,
	}
	return &command
}

// Run commands without any output
func (c *Commands) Run(command string, args ...string) {
	args = append([]string{"-c", command}, args...)
	cmd := exec.Command("/bin/sh", args...)
	if err := cmd.Run(); err != nil {
		c.Error = err
		log.Fatal(err)
	}
}

// Run outer commands
func (c *Commands) OuterRun(command string, args ...string) {
	_, err := exec.Command("tesseract", args...).Output()
	if err != nil {
		c.Error = err
		log.Fatal(err)
	}
}

// GetOutput run commands without status, but the output
func (c *Commands) GetOutput(command string, args ...string) (output string) {
	args = append([]string{"-c", command}, args...)
	bytes, err := exec.Command("/bin/sh", args...).Output()
	if err != nil {
		c.Output = err.Error()
		c.Error = err
	} else {
		c.Output = strings.Trim(string(bytes), "\n")
	}
	return c.Output
}

// GetStatusOutput run commands with status and output
func (c *Commands) GetStatusOutput(command string, args ...string) (status bool, output string) {
	args = append([]string{"-c", command}, args...)
	cmd := exec.Command("/bin/sh", args...)
	var bufferout, bufferin, buffererr bytes.Buffer
	cmd.Stdout = &bufferout
	cmd.Stdin = &bufferin
	cmd.Stderr = &buffererr

	// do the commands
	err := cmd.Run()
	if err != nil {
		c.Output = err.Error()
		c.Error = err
	} else {
		c.Status = cmd.ProcessState.Success()
		c.Output = bufferout.String()
	}
	return c.Status, c.Output
}

// just for debugging
func (c *Commands) GetError() error {
	return c.Error
}