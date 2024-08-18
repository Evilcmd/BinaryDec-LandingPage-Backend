package server

import (
	"log"
	"net/http"

	"github.com/Evilcmd/Hackup-backend/internal/apis"
)

func NewServer(port string, apiCfg *apis.ApiConfig) *http.Server {

	if port == "" {
		log.Fatal("port not found")
	}
	server := &http.Server{
		Addr:    ":" + port,
		Handler: GetRouter(apiCfg),
	}

	return server
}
