package config

import (
	"context"
)

type Config struct {
	PerSecond int
	SizeFrom  int
	SizeTo    int
	Limit     int
}

func New(ctx context.Context) (*Config, error) {
	var cfg Config

	var err error
	cfg.PerSecond, cfg.SizeFrom, cfg.SizeTo, cfg.Limit, err = getEnv()
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
