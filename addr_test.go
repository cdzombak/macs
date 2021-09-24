package macs

import (
	"encoding/json"
	"net"
	"testing"
)

func TestParseMAC(t *testing.T) {
	mac, err := ParseMAC("78:4f:43:87:9e:f4")
	if err != nil {
		t.Error(err)
	}
	if mac.String() != "78:4f:43:87:9e:f4" {
		t.Errorf("mac is wrong: %s", mac.String())
	}

	mac, err = ParseMAC("foo:bar")
	if err == nil || mac != nil {
		t.Error("ParseMAC should have failed")
	}
}

func TestMACMarshalJson(t *testing.T) {
	mac := MustParseMAC("78:4f:43:87:9e:f4")
	jsonVal, err := json.Marshal(mac)
	if err != nil {
		t.Error(err)
	}

	if string(jsonVal) != "\"78:4f:43:87:9e:f4\"" {
		t.Errorf("json is wrong: %s", string(jsonVal))
	}
}

func TestMACUnmarshalJSON(t *testing.T) {
	mac := Addr{}
	err := json.Unmarshal([]byte("\"78:4f:43:87:9e:f4\""), &mac)
	if err != nil {
		t.Error(err)
	}
	if mac.String() != "78:4f:43:87:9e:f4" {
		t.Errorf("unmarshaled mac is wrong: %s", mac.String())
	}
}

func TestZeroValue(t *testing.T) {
	mac := Addr{}
	if !mac.IsZeroValue() {
		t.Error("an empty mac is its zero value")
	}
}

func TestString(t *testing.T) {
	mac := MustParseMAC("78:4F:43:87:9E:F4")
	if mac.String() != "78:4f:43:87:9e:f4" {
		t.Errorf("stringed mac is wrong: %s", mac.String())
	}
}

func TestNetHardwareAddrCompat(t *testing.T) {
	netAddr, err := net.ParseMAC("78:4f:43:87:9e:f4")
	if err != nil {
		t.Fatal("net.ParseMAC must work")
	}

	mac := MustParseMAC("78:4f:43:87:9e:f4")
	if mac.NetHardwareAddr().String() != netAddr.String() {
		t.Error("net.Addr should be the same")
	}

	macFromNetAddr := HardwareAddrFromNet(netAddr)
	if macFromNetAddr.String() != mac.String() {
		t.Error("net.Addr should be the same")
	}
}
