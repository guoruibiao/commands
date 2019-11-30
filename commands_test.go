package commands

import (
	"fmt"
	"testing"
)

func Test_Run(t *testing.T) {
	c := New()
	c.Run("ls", "-alhtr")

	c.Run("du", "-alh")
}

func Test_GetOutput(t *testing.T) {
	c := New()
	output := c.GetOutput("ls", "-al")
	fmt.Println(output)

	output = c.GetOutput("ls", "-al")
	fmt.Println(output)

	output = c.GetOutput("ls -alhtr|head -2")
	t.Log(output)
}

func Test_GetStatusOutput(t *testing.T) {
	c := New()
	status, output := c.GetStatusOutput("ls  -alhtr")
	if status == false {
		t.Error(output)
	} else {
		t.Log(output)
	}
	t.Log(c.GetError())
	t.Log( c.Output)
}
