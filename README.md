# Commands 
<<<<<<< HEAD
=======

[![GoDoc](https://godoc.org/github.com/guoruibiao/commands?status.svg)](https://godoc.org/github.com/guoruibiao/commands)



>>>>>>> c267f0f355fd99f015dd5925a4dfd945363656bf
reference from `commands` in Python. Do some commands if you need, with optional choice of `Status` and `Output`.

## Support features

- `Run(command string, args...string)`  without any output
- `GetOutput(command string, args...string) (output string)` with output of command
- `GetStatusOutput(command string, args...string) (status bool, output string)` with status and output of commands


## Demos
```
package main

import (
	"fmt"
	"runtime"

	"github.com/guoruibiao/commands"
)

func main() {
	c := commands.New()
	c.Run("ls", "-ltrha")
	output := c.GetOutput("ls", "-ltrha")
	fmt.Println(output)
	status, output := c.GetStatusOutput("du", "-ah")
	fmt.Println("status: ", status)
	fmt.Println("output: ", output)

	// macOS上弹出一个notification
	if runtime.GOOS == "darwin" {
		c.Run("osascript", "-e", `display notification "内容体" with title "标题"`)
	} else {
		fmt.Println("其他操作系统弹出Notification的示例")
	}
}

```

There are also some new features to be added, so happy to see your PRs.

<<<<<<< HEAD
finally, enjoy it.
=======
finally, enjoy it.
>>>>>>> c267f0f355fd99f015dd5925a4dfd945363656bf
