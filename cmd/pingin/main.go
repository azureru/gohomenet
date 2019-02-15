package main

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

// ENV
//    PINGIN_INTERFACE=en0
//    PINGIN_INTERFACE=en0,en6
func main() {
	log.SetFlags(log.Lshortfile | log.LUTC)
	interfaceString := os.Getenv("PINGIN_INTERFACE")
	interfaceNames := strings.Split(interfaceString, ",")

	// if interface not specified use all :)
	if len(interfaceNames) <= 0 || interfaceString == "" {
		ifns, err := net.Interfaces()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		for _, ifn := range ifns {
			log.Println(ifn.Name)
			interfaceNames = append(interfaceNames, ifn.Name)
		}
	}

	for _, name := range interfaceNames {
		err := getIP(name)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func getIP(interfaceName string) error {
	ief, err := net.InterfaceByName(interfaceName)
	if err != nil {
		return err
	}
	addrs, err := ief.Addrs()
	if err != nil {
		return err
	}
	var ip net.IP
	for _, addr := range addrs {
		// use non IPV6 addr
		if !strings.Contains(addr.String(), ":") {
			if !strings.HasPrefix(addr.String(), "127.0.0") {
				ip = addr.(*net.IPNet).IP
			}
		}
	}
	tcpAddr := &net.TCPAddr{
		IP: ip,
	}
	transport := &http.Transport{
		Proxy:               http.ProxyFromEnvironment,
		Dial:                (&net.Dialer{LocalAddr: tcpAddr}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
	}
	client := &http.Client{
		Transport: transport,
	}
	response, err := client.Get("https://api.ipify.org/")
	if err != nil {
		return err
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}
		var contentsStr = string(contents)
		log.Printf("%s %s\n", interfaceName, contentsStr)
		return nil
	}
}
