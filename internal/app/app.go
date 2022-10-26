package app

import (
	"context"

	"github.com/zvlb/log-spammer/pkg/config"
	"github.com/zvlb/log-spammer/pkg/spamer"
)

// type App struct{}

func NewApp(ctx context.Context) error {
	// Create gracefull shutdown

	// Read params
	cfg, err := config.New(ctx)
	if err != nil {
		return err
	}

	spamer.Spamer(*cfg)

	return nil
}
