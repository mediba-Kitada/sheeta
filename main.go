package main

import (
	"flag"
	"fmt"
	"os"

	"context"

	"github.com/google/subcommands"
	"github.com/mediba-Kitada/sheeta/commands"
)

// Version of Sheeta
var (
	Version = "0.1.0"
)

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&commands.StartCmd{}, "start")

	var v = flag.Bool("v", false, "Show version")

	flag.Parse()

	if *v {
		fmt.Printf("sheeta %s\n", version)
		os.Exit(int(subcommands.ExitSuccess))
	}

	ctx := context.Background()

	os.Exit(int(subcommands.Execute(ctx)))
}
