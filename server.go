package main

import (
	"github.com/acm-uiuc/arbor"
	"github.com/acm-uiuc/groot-api-gateway/config"
	"github.com/acm-uiuc/groot-api-gateway/services"
)

func main() {
	config.LoadArborConfig()
	Routes := services.RegisterAPIs()
	arbor.Boot(Routes, 8000)
}
