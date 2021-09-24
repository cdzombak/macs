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

func ParseMAC(mac string) (*Addr, error) {
	if addr, err := net.ParseMAC(mac); err != nil {
		return nil, err
	} else {
		return &Addr{addr: addr}, nil
	}
}

// MustParseMAC parses the given string to a net.HardwareAddr, or panics if the
// string doesn't represent a valid MAC address.
func MustParseMAC(mac string) *Addr {
	if retv, err := ParseMAC(mac); err != nil {
		panic(err)
	} else {
		return retv
	}
}

func (a *Addr) IsZeroValue() bool {
	return len(a.addr) == 0
}

func (a *Addr) String() string {
	return a.addr.String()
}

func (a *Addr) Equals(other *Addr) bool {
	return bytes.Compare(a.addr, other.addr) == 0
}

func (a *Addr) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.addr.String())
}

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

func (a *Addr) NetHardwareAddr() net.HardwareAddr {
	return a.addr
}

func AddrFromNet(addr net.HardwareAddr) *Addr {
	return &Addr{addr: addr}
}
