package main

import (
	"fmt"
	"os"

	"github.com/pengwynn/flint/flint"
)

func main() {
	app := flint.NewApp()

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
        PrintoutSystemln("first_comment");
		os.Exit(1)
	}
}
