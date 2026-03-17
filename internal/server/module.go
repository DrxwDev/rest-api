// Package server
package server

import (
	"net/http"

	"go.uber.org/fx"
)

var Module = fx.Module("server",
	fx.Provide(
		NewGinRouter,
		NewHTTPServer,
	),
	fx.Invoke(
		Routes,
		func(_ *http.Server) {},
	),
)
