// Package options contains program options.
package options

import "flag"

// Options contains configurable options for command flags.
type Options struct {
	Port uint
	Help bool
}

// BindFlags binds command flags to an options instance using the
// provided flagset.
func (env *Options) BindFlags(fs *flag.FlagSet) {
	fs.Usage = func() {
		fs.PrintDefaults()
	}

	fs.BoolVar(&env.Help, "help", false, "Print this help message.")
	fs.BoolVar(&env.Help, "h", false, "Print this help message.")
	fs.UintVar(&env.Port, "port", 3202, "Port to listen on.")
}
