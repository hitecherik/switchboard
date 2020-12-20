package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/hitecherik/switchboard/internal/router"
	"github.com/hitecherik/switchboard/internal/walker"
	"github.com/pelletier/go-toml"
)

var (
	temporary bool
	port      uint
	tree      *toml.Tree
)

func bail(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func init() {
	var path string

	flag.BoolVar(&temporary, "temporary", false, "whether the redirect should be temporary")
	flag.UintVar(&port, "port", 8000, "port to listen on")
	flag.StringVar(&path, "config", "config.toml", "file defining redirects")
	flag.Parse()

	var err error
	tree, err = toml.LoadFile(path)
	bail(err)
}

func main() {
	redirects, err := walker.Walk(tree)
	bail(err)

	http.Handle("/", router.BuildRouter(redirects, temporary))
	log.Printf("Listening on port %v...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
