package main

import (
	"net/http"

	"github.com/caarlos0/env/v11"
	"github.com/gleich/lumber/v3"
)

type config struct {
	URL  *string `env:"REDIRECT_URL"`
	CODE *int    `env:"REDIRECT_STATUS_CODE"`
}

func main() {
	lumber.Info("booted")

	conf := loadConfig()

	mux := http.NewServeMux()
	mux.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, *conf.URL, *conf.CODE)
			lumber.Done("Redirected", r.RequestURI, "to", *conf.URL)
		},
	)

	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		lumber.Fatal(err, "starting server failed")
	}
}

func loadConfig() config {
	conf, err := env.ParseAs[config]()
	if err != nil {
		lumber.Fatal(err, "loading .env file failed")
	}
	if conf.URL == nil {
		lumber.Fatal(err, "REDIRECT_URL env var is not set.")
	}
	if conf.CODE == nil {
		defaultCode := http.StatusTemporaryRedirect
		conf.CODE = &defaultCode
	}
	return conf
}
