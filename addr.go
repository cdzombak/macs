package macs

import (
	"bytes"
	"encoding/json"
	"net"
)

// Addr is a hardware/MAC address which supports JSON [un]marshaling.
type Addr struct {
	addr net.HardwareAddr
}

// ParseMAC parses the given string to an Addr, returning an error if parsing fails.
func ParseMAC(mac string) (*Addr, error) {
	if addr, err := net.ParseMAC(mac); err != nil {
		return nil, err
	} else {
		return &Addr{addr: addr}, nil
	}
}

// MustParseMAC parses the given string to an Addr, or panics if parsing fails.
func MustParseMAC(mac string) *Addr {
	if retv, err := ParseMAC(mac); err != nil {
		panic(err)
	} else {
		return retv
	}
}

// IsZeroValue returns true if the MAC address is the zero value.
func (a *Addr) IsZeroValue() bool {
	return len(a.addr) == 0
}

// String returns the string representation of the MAC address.
func (a *Addr) String() string {
	return a.addr.String()
}

// Equals returns true if the MAC address is equal to the given MAC address.
func (a *Addr) Equals(other *Addr) bool {
	return bytes.Compare(a.addr, other.addr) == 0
}

// MarshalJSON marshals the MAC address to JSON.
func (a *Addr) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.addr.String())
}

// UnmarshalJSON unmarshals the MAC address from JSON.
func (a *Addr) UnmarshalJSON(bytes []byte) error {
	macStr := ""
	if err := json.Unmarshal(bytes, &macStr); err != nil {
		return err
	}
	if addr, err := net.ParseMAC(macStr); err != nil {
		return err
	} else {
		a.addr = addr
		return nil
	}
}

// NetHardwareAddr returns the MAC address as a net.HardwareAddr.
func (a *Addr) NetHardwareAddr() net.HardwareAddr {
	return a.addr
}

// AddrFromNet creates an Addr from a net.HardwareAddr.
func AddrFromNet(addr net.HardwareAddr) *Addr {
	return &Addr{addr: addr}
}
