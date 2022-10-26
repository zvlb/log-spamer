package config

import (
	"context"
)

type Config struct {
	PerSecond int
	SizeFrom  int
	SizeTo    int
	ctx       context.Context
}

func New(ctx context.Context) (*Config, error) {
	var cfg Config

	var err error
	cfg.PerSecond, cfg.SizeFrom, cfg.SizeTo, err = getEnv()
	if err != nil {
		return nil, err
	}

	cfg.ctx = ctx

	return &cfg, nil
}
func (c *Config) Context() context.Context {
	return c.ctx
}
