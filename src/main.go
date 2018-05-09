package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/ContinuumLLC/boaring/src/common"
	"github.com/ContinuumLLC/boaring/src/service"
	"github.com/fatih/color"
	"github.com/urfave/cli"
)

func main() {
	//testColor()
	testCLI()
}

func testCLI() {
	app := cli.NewApp()
	app.Name = "greet"
	app.Usage = "fight the loneliness!"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Hello friend!")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func testColor() {
	red := color.New(color.FgRed).PrintlnFunc()
	red("Warning")
	err := errors.New("Hello Boaring")
	red("Error: %s", err)
	common.Hello()
	service.Hello()
}
