// Package database
package database

import "go.uber.org/fx"

var Module = fx.Module("database",
	fx.Provide(
		NewPool,
	),
)
