package main

import (
	"fmt"

	"github.com/u-root/webboot/pkg/wifi"
)

// Interface is for network interfaces.
type Interface struct {
	label string
}

// Label returns a label (e.g. wlan0) for an Interface.
func (i *Interface) Label() string {
	return i.label
}

// Network is a WIFI network.
type Network struct {
	info wifi.Option
}

// Label returns information about a Network.
func (n *Network) Label() string {
	switch n.info.AuthSuite {
	case wifi.NoEnc:
		return fmt.Sprintf("%s: No Passphrase\n", n.info.Essid)
	case wifi.WpaPsk:
		return fmt.Sprintf("%s: WPA-PSK (only passphrase)\n", n.info.Essid)
	case wifi.WpaEap:
		return fmt.Sprintf("%s: WPA-EAP (passphrase and identity)\n", n.info.Essid)
	case wifi.NotSupportedProto:
		return fmt.Sprintf("%s: Not a supported protocol\n", n.info.Essid)
	}
	return "Invalid wifi network."
}
