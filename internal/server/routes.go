package server

import (
	"net/http"

	"github.com/Evilcmd/Hackup-backend/internal/apis"
	"github.com/Evilcmd/Hackup-backend/internal/middleware"
)

func GetRouter(apiCfg *apis.ApiConfig) http.Handler {
	router := http.NewServeMux()

	generalRouter := generalRoutes(apiCfg)

	router.Handle("/", generalRouter)

	return middleware.Cors(router)
}

func generalRoutes(apiCfg *apis.ApiConfig) http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("/", apis.Root)
	router.HandleFunc("GET /health", apis.CheckHealth)
	router.HandleFunc("GET /err", apis.ErrCheck)

	router.HandleFunc("POST /presignup", apiCfg.Presignup)

	return router
}
