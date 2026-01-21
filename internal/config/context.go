package config

import "context"

type contextKey struct{}

func WithConfig(ctx context.Context, cfg *Config) context.Context {
	return context.WithValue(ctx, contextKey{}, cfg)
}

func FromContext(ctx context.Context) (*Config, bool) {
	cfg, ok := ctx.Value(contextKey{}).(*Config)
	return cfg, ok
}
