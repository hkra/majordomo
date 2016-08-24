package main

import (
	"flag"
	"log"
)

// Options is a container for environment values and program parameters.
type Options struct {
	Port uint
	Help bool
}

var env = new(Options)

func init() {
	flag.BoolVar(&env.Help, "help", false, "Print this help message.")
	flag.BoolVar(&env.Help, "h", false, "Print this help message.")
	flag.UintVar(&env.Port, "port", 3202, "Port to listen on.")
	flag.Parse()
}

func main() {
	if env.Help {
		flag.Usage()
		return
	}

	log.Printf("Starting majordomo command broker on port [%d]", env.Port)
}
