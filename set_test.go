package macs

import (
	"encoding/json"
	"testing"
)

func TestEmptyMACSet(t *testing.T) {
	set := EmptySet()
	if set.Len() != 0 {
		t.Error("empty set should have length 0")
	}
}

func TestMACSetBasics(t *testing.T) {
	set := EmptySet()
	set.Add(MustParseMAC("78:4f:43:87:9e:f4"))
	set.Add(MustParseMAC("78:4F:43:87:9E:F4"))
	set.Add(MustParseMAC("de:ad:be:ef:aa:bb"))
	set.Add(MustParseMAC("66:66:66:66:66:66"))

	if set.Len() != 3 {
		t.Error("set must have 3 elements")
	}
	if !set.Contains(MustParseMAC("78:4F:43:87:9E:F4")) {
		t.Error("set is missing something we added")
	}
	if set.Contains(MustParseMAC("00:11:22:33:22:11")) {
		t.Error("set has something we did not add")
	}

	set.Remove(MustParseMAC("66:66:66:66:66:66"))
	if set.Len() != 2 {
		t.Error("set must have 3 elements")
	}
	if set.Contains(MustParseMAC("66:66:66:66:66:66")) {
		t.Error("set has something we removed")
	}
}

func TestMACSetAll(t *testing.T) {
	set := EmptySet()
	set.Add(MustParseMAC("78:4F:43:87:9E:F4"))
	set.Add(MustParseMAC("de:ad:be:ef:aa:bb"))
	set.Add(MustParseMAC("66:66:66:66:66:66"))

	for _, mac := range set.All() {
		if mac.String() != "78:4f:43:87:9e:f4" && mac.String() != "de:ad:be:ef:aa:bb" && mac.String() != "66:66:66:66:66:66" {
			t.Error("set did not produce what we put in")
		}
	}
}

func TestMACSetAddAll(t *testing.T) {
	set := EmptySet()
	set.Add(MustParseMAC("78:4F:43:87:9E:F4"))
	set.Add(MustParseMAC("de:ad:be:ef:aa:bb"))
	set.Add(MustParseMAC("66:66:66:66:66:66"))

	set2 := EmptySet()
	set2.Add(MustParseMAC("00:11:22:33:22:11"))
	set2.AddAllFrom(set)

	if set2.Len() != 4 {
		t.Error("set2 does not have the right length")
	}
	if !set2.Contains(MustParseMAC("de:ad:be:ef:aa:bb")) {
		t.Error("set2 does not have something that should have been added from set1")
	}
}

func TestMACSetJSON(t *testing.T) {
	set := EmptySet()
	set.Add(MustParseMAC("78:4F:43:87:9E:F4"))
	set.Add(MustParseMAC("de:ad:be:ef:aa:bb"))
	set.Add(MustParseMAC("66:66:66:66:66:66"))

	jsonBytes, err := set.MarshalJSON()
	if err != nil {
		t.Error("marshaling must not fail")
	}
	if len(jsonBytes) == 0 {
		t.Error("marshaling must produce data")
	}

	var set2 *Set
	err = json.Unmarshal(jsonBytes, &set2)
	if err != nil {
		t.Error("unmarshal must not fail")
	}
	if set2.Len() != 3 {
		t.Error("set2 must have length 3")
	}
	for _, mac := range set2.All() {
		if mac.String() != "78:4f:43:87:9e:f4" && mac.String() != "de:ad:be:ef:aa:bb" && mac.String() != "66:66:66:66:66:66" {
			t.Error("unmarshaled set did not produce what we put in")
		}
	}

	if set2.Contains(MustParseMAC("00:11:22:33:22:11")) {
		t.Error("unmarshaled set must not have something we didn't put in")
	}
}

func TestMACSetEmptyJSON(t *testing.T) {
	set := EmptySet()
	jsonBytes, err := set.MarshalJSON()
	if err != nil {
		t.Error("marshaling must not fail")
	}
	if len(jsonBytes) != 2 {
		t.Error("marshaling must produce an empty list")
	}
	var set2 *Set
	err = json.Unmarshal(jsonBytes, &set2)
	if err != nil {
		t.Error("unmarshal must not fail")
	}
	if set2.Len() != 0 {
		t.Error("set2 must have length 0")
	}
}
