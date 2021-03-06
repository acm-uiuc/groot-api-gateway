package main

import (
	"github.com/acm-uiuc/groot-api-gateway/config"
	"github.com/acm-uiuc/groot-api-gateway/services"
	"github.com/arbor-dev/arbor"
)

func main() {
	config.LoadArborConfig()
	Routes := services.RegisterAPIs()
	arbor.Boot(Routes, "0.0.0.0", 8000)
}
