package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
)

var defaultEnv = map[string]string{
	"SERVER_HOSTNAME": "https://mozilla.cloudflare-dns.com/dns-query",
	"DEBUG":           "FALSE",
}

func getEnv(env string) string {
	newEnv := os.Getenv(env)
	if newEnv != "" {
		return newEnv
	}
	if v, ok := defaultEnv[env]; ok {
		return v
	}
	return ""
}

func main() {
	debug := getEnv("DEBUG")
	dohHost := getEnv("SERVER_HOSTNAME")

	if debug == "TRUE" {
		log.SetFlags(log.Lshortfile)
	} else {
		log.SetFlags(0)
	}

	err := newUDPServer(dohHost)
	if err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}

func newUDPServer(dohHost string) error {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP(dohHost), Port: 53})
	if err != nil {
		return err
	}
	for {
		var raw [512]byte
		n, addr, err := conn.ReadFromUDP(raw[:512])
		if err != nil {
			log.Printf("could not read: %s", err)
			continue
		}
		log.Printf("new connection from %s:%d", addr.IP.String(), addr.Port)
		go proxy(dohHost, conn, addr, raw[:n])
	}
}

func proxy(dohHost string, conn *net.UDPConn, addr *net.UDPAddr, raw []byte) {
	enc := base64.RawURLEncoding.EncodeToString(raw)
	url := fmt.Sprintf("%s?dns=%s", dohHost, enc)
	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Printf("could not create request: %s", err)
		return
	}
	r.Header.Set("Content-Type", "application/dns-message")
	r.Header.Set("Accept", "application/dns-message")

	c := http.Client{}
	resp, err := c.Do(r)
	if err != nil {
		log.Printf("could not perform request: %s", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("wrong response from DOH server got %s", http.StatusText(resp.StatusCode))
		return
	}

	msg, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("could not read message from response: %s", err)
		return
	}

	if _, err := conn.WriteToUDP(msg, addr); err != nil {
		log.Printf("could not write to udp connection: %s", err)
		return
	}
}
