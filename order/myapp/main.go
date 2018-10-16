package main

import (
	"myapp/http"
)
func main() {

	go http.Start()

	select {}
}
