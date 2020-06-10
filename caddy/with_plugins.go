package main

import (
	//"github.com/mholt/caddy/caddy/caddymain"
	"github.com/GuoxiW/caddy/caddy/caddymain"

	// plug in plugins here
	//_ "github.com/captncraig/cors/caddy"
	_ "github.com/GuoxiW/cors/caddy"
)

func main() {
	// optional: disable telemetry
	caddymain.EnableTelemetry = false
	caddymain.Run()
}
