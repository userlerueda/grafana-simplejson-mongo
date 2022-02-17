package main

import (
	"log"
	"os"
	"strconv"

	"github.com/m-masataka/grafana-simplejson-mongo/api"
)

func main() {
	p, err := strconv.Atoi(os.Getenv("GSM_PORT"))
	if err != nil {
		p = 8080
	}
	m := os.Getenv("GMS_MONGOHOST")
	if m == "" {
		m = "localhost"
	}
	conf := api.Config{
		Port:      p,
		MongoHost: m,
	}
	errs := make(chan error, 2)
	api.StartHTTPServer(conf, errs)
	log.Println("start")
	for {
		err := <-errs
		log.Println(err)
	}
}
