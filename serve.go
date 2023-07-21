package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"strconv"
)

func main() {
	port := flag.Int("port", 3000, "Port to run web server to")
	flag.Parse()

	ip, err := getIPAddress()
	if err != nil {
		fmt.Println(err)
		return
	}

	addr := ip + ":" + strconv.Itoa(*port)
	fmt.Println("Listening on http://" + addr)
	http.ListenAndServe(addr, http.FileServer(http.Dir(".")))
}

func getIPAddress() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			return ipnet.IP.String(), nil
		}
	}
	return "", fmt.Errorf("Unable to determine IP address")
}
