// Package options contains program options.
package options

import "flag"

// Options contains configurable options for command flags.
type Options struct {
	Port        string
	BindAddress string
	UseTLS      bool
	Help        bool
}

// BindFlags binds command flags to an options instance using the
// provided flagset.
func (env *Options) BindFlags(fs *flag.FlagSet) {
	fs.Usage = func() {
		fs.PrintDefaults()
	}

	fs.BoolVar(&env.Help, "help", false, "Print this help message.")
	fs.BoolVar(&env.Help, "h", false, "Print this help message.")
	fs.StringVar(&env.Port, "port", "3202", "Port to listen on.")
	fs.StringVar(&env.BindAddress, "address", "localhost", "host address to which to bind.")
	fs.BoolVar(&env.UseTLS, "useTls", false, "True to use TLS. Certificates should be provided when enabled.")
}
