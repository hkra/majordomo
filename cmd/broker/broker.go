package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/hkra/majordomo/cmd/broker/app"
	"github.com/hkra/majordomo/cmd/broker/app/options"
)

func parseProgramFlags(flags *flag.FlagSet) *options.Options {
	var env = options.Options{}
	env.BindFlags(flags)

	args := os.Args[1:]
	if err := flags.Parse(args); err != nil {
		log.Fatal("Flag parsing failed", args)
	}

	if env.Help {
		fmt.Printf("Usage: %s\n", os.Args[0])
		flags.Usage()
		os.Exit(0)
	}

	return &env
}

func main() {
	flags := parseProgramFlags(flag.NewFlagSet("broker", flag.ExitOnError))
	server := app.New(flags)
	server.Start()
}
