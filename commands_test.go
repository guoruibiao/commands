package commands

import (
	"fmt"
	"testing"
)

func Test_Run(t *testing.T) {
	c := New()
	c.Run("ls", "-alhtr")
}

func Test_GetOutput(t *testing.T) {
	c := New()
	output := c.GetOutput("ls", "-al")
	fmt.Println(output)
}

func Test_GetStatusOutput(t *testing.T) {
	c := New()
	status, output := c.GetStatusOutput("ls", "-alhtr")
	if status == false {
		t.Error(output)
	} else {
		fmt.Println(output)
	}

}
