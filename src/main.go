package main

import (
	"errors"

	"github.com/ContinuumLLC/boaring/src/common"
	"github.com/ContinuumLLC/boaring/src/service"
	"github.com/fatih/color"
)

func main() {
	red := color.New(color.FgRed).PrintlnFunc()
	red("Warning")
	err := errors.New("Hello Boaring")
	red("Error: %s", err)
	common.Hello()
	service.Hello()
}
