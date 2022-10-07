package cmd

import (
	"docker-golang/config"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

var (
	env = config.Env
)

type Server struct{}

func (s *Server) Start() {
	// * Configure *
	log.Info().Msg("Creating configs...")
	config.Configure()

	// * Http *
	log.Info().Msg("Creating http server...")
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Olá mundo\n")
	})

	http.ListenAndServe(fmt.Sprintf(":%s", env.Port), nil)
}
