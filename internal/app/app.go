package app

import (
	"log"

	"github.com/knadh/koanf/v2"
)

func Run(cfg *koanf.Koanf) {
	host := cfg.String("http.host")
	port := cfg.String("http.port")

	log.Printf("Serving HTTP on %v port %v ...\n", host, port)
}
