package main

import (
	"net"
	"os"

	"github.com/charmbracelet/log"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Error("could not get network interfaces", err)
		os.Exit(0)
	}

	log.Info("found interfaces...")

	for _, iface := range interfaces {
		addresses, err := iface.Addrs()
		if err != nil {
			log.Error("cannot read address", err)
		}

		for _, address := range addresses {
			if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {

				if ipnet.IP.To4() != nil {
					log.Info("address::", iface.HardwareAddr.String())
					log.Info("network ip:", ipnet.IP.To4())
				}
			}
		}
	}

	log.Info("completed.")
}
