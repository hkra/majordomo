package app

import (
	"testing"

	"github.com/hkra/majordomo/cmd/broker/app/options"
)

func configTransfersSettingsFromOptions(t *testing.T) {
	options := options.Options{
		Port: 12345,
	}
	server := APIServer{}
	server.Config(&options)
	if server.Port != 12345 {
		t.Fail()
	}
}
