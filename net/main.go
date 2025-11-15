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
		log.Info("interface name: ", iface.Name)
		log.Info("interface addres: ", iface.HardwareAddr)

		addresses, err := iface.Addrs()
		if err != nil {
			log.Error("cannot read address", err)
		}

		for _, address := range addresses {
			log.Info("network address:", address.String())
		}
		log.Info("\n")
	}
}
