package main

import "goRank/webserver"
import (
	"goRank/storage/elastic"
)

func main() {
	se := elastic.Elastic{}
	se.Connect()
	ws := webserver.WebServer{}
	ws.Serve(se)
}
