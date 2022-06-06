package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetConfig(t *testing.T) {
	req := require.New(t)
	os.Chdir("../../")

	t.Run("Not nil config", func(t *testing.T) {
		cfg, _ := GetConfig()
		req.NotNil(cfg)
	})

	t.Run("Check app listen port", func(t *testing.T) {
		cfg, _ := GetConfig()
		req.Equal("8000", cfg.Listen.Port)
	})
}

