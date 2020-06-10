package main

import (
	//"github.com/mholt/caddy/caddy/caddymain"
	"github.com/GuoxiW/caddy/tree/master/caddy/caddymain"

	// plug in plugins here
	//_ "github.com/captncraig/cors/caddy"
	_ "/github.com/GuoxiW/cors/tree/master/caddy"
)

func main() {
	// optional: disable telemetry
	caddymain.EnableTelemetry = false
	caddymain.Run()
}
