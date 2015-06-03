package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gernest/aurora"
)

func main() {
	d, err := ioutil.ReadFile("config/app/app.json")
	if err != nil {
		panic(err)
	}
	cfg := &aurora.RemixConfig{}
	err = json.Unmarshal(d, cfg)
	if err != nil {
		panic(err)
	}
	rx := aurora.NewRemix(cfg)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))
	http.Handle("/", rx.Routes())
	port := os.Getenv("PORT")
	log.Println("starting at port ", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
