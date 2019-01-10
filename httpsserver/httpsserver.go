package main

import (
        "fmt"
        "net"
        "net/http"
        "os"
)

var ip string

func GetLocalIP() {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
		fmt.Println("ERROR:" + err.Error())
		os.Exit(1)
	}
    for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
                ip = ipnet.IP.String()
			}
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Server IP: <font color=blue>%s</font>, Client IP: %s.</h1>", ip, r.RemoteAddr)
}

func main() {
    GetLocalIP()
    fmt.Println(ip)
	http.HandleFunc("/", handler)
    http.ListenAndServeTLS("0.0.0.0:9420", "server.crt", "server.key", nil)
}

