package app

import (
	"testing"

	"github.com/hkra/majordomo/cmd/broker/app/options"
)

func configTransfersSettingsFromOptions(t *testing.T) {
	config := configureAPIServer(&options.Options{
		Port:        "12345",
		BindAddress: "localhost",
		UseTLS:      true,
	})

	if config.Port != "12345" ||
		config.BindAddress != "localhost" ||
		config.UseTLS != true {
		t.Fail()
	}
}
