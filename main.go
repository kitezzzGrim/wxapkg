package main

import (
	"github.com/fatih/color"
	"os"
	"wxapkg/cmd"
)

var (
	err_  = color.New(color.FgRed)
	warn_ = color.New(color.FgYellow)
	info_ = color.New(color.FgCyan)
)

func main() {
	if err := cmd.New().Run(os.Args); err != nil {
		err_.Println()
		err_.Printf("[-]  %s\n", err.Error())
	}
}
