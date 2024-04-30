package main

import (
	"fmt"
	"github.com/titancodehub/invoicekit/subcommand"
	"github.com/titancodehub/invoicekit/tui/styles"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Print(styles.InfoTextStyle.Render("Usage: invk <subcommand> [<args>]\n"))
		os.Exit(1)
	}

	switch os.Args[1] {
	case subcommand.InitCommandName:
		subcommand.Initialization()
		break
	}

}
