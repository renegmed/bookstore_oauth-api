package main

import (
	"bookstore_oauth-api/src/app"
	"flag"
	"log"
)

// var usageStr = `
// Usage: Oauth API

// Options:
// 	-a,  --addr  Server URL address e.g. localhost:8080
// `

// func usage() {
// 	log.Printf("%s\n", usageStr)
// 	os.Exit(0)
// }

func main() {

	var (
		URL string
	)

	// flag.StringVar(&URL, "s", "localhost:8080", "server URL address")
	flag.StringVar(&URL, "server", "localhost:8080", "server URL address")

	// flag.Usage = usage
	flag.Parse()

	log.Println("... server", URL)

	app.StartApplication(URL)
}
