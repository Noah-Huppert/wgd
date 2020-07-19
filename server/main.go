package main

import (
	"net/http"

	"github.com/Noah-Huppert/goconf"
	"github.com/Noah-Huppert/golog"
	"github.com/gin-gonic/gin"
	"github.com/vishvananda/netlink"
)

// Application config.
type Config struct {
	// TODO keep defining
}

func main() {
	// Logger
	logger := golog.NewLogger("wgd.server")
	logger.Debug("Starting")

	// Load configuration
	cfgLdr := goconf.NewLoader()
	cfgLdr.AddConfigPath("/etc/wgd/server.d/*")
	config := Config{}
	if err := cfgLdr.Load(&config); err != nil {
		return nil, fmt.Errorf("Failed to load configuration: %s", err)
	}
}
