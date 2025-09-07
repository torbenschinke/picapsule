package settings

import (
	"net"
	"strings"
)

func GetNetworks() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return err.Error()
	}

	var networks []string
	for _, i := range interfaces {
		addrs, err := i.Addrs()
		if err != nil {
			networks = append(networks, err.Error())
			continue
		}

		for _, addr := range addrs {
			networks = append(networks, addr.String())
		}
	}

	return strings.Join(networks, "\n")
}
