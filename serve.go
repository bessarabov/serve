package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
)

func main() {

	port := flag.Int("port", 3000, "Port to run web server to")
	flag.Parse()

	addr := "127.0.0.1:" + strconv.Itoa(*port)
	fmt.Println("Listening on http://" + addr)
	http.ListenAndServe(addr, http.FileServer(http.Dir(".")))
}
