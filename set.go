package macs

import (
	"encoding/json"
)

// Set is a (non-concurrency-safe!) set of MAC addresses which supports [un]marshaling
// to/from a JSON array.
type Set struct {
	macs map[string]bool
}

// MarshalJSON marshals the set to a JSON array of strings.
func (s *Set) MarshalJSON() ([]byte, error) {
	keys := make([]string, 0, len(s.macs))
	for k, _ := range s.macs {
		keys = append(keys, k)
	}
	return json.Marshal(keys)
}

// UnmarshalJSON unmarshals a JSON array of strings into the set.
func (s *Set) UnmarshalJSON(bytes []byte) error {
	var macsList []string
	if err := json.Unmarshal(bytes, &macsList); err != nil {
		return err
	}
	for _, macStr := range macsList {
		if mac, err := ParseMAC(macStr); err != nil {
			return err
		} else {
			s.Add(mac)
		}
	}
	return nil
}

func (s *Set) ensureInit() {
	if s.macs == nil {
		s.macs = make(map[string]bool)
	}
}

// Contains returns true if the set contains the given MAC address.
func (s *Set) Contains(mac *Addr) bool {
	s.ensureInit()
	return s.macs[mac.String()]
}

// Add adds the given MAC address to the set.
func (s *Set) Add(mac *Addr) {
	s.ensureInit()
	s.macs[mac.String()] = true
}

// Remove removes the given MAC address from the set.
func (s *Set) Remove(mac *Addr) {
	s.ensureInit()
	delete(s.macs, mac.String())
}

// All returns a slice of all MAC addresses in the set.
func (s *Set) All() []*Addr {
	s.ensureInit()
	retv := make([]*Addr, len(s.macs))
	i := 0
	for k := range s.macs {
		retv[i] = MustParseMAC(k)
		i++
	}
	return retv
}

// Len returns the number of MAC addresses in the set.
func (s *Set) Len() int {
	return len(s.macs)
}

// AddAllFrom adds all MAC addresses from the given set to this set.
func (s *Set) AddAllFrom(other *Set) {
	s.ensureInit()
	for _, mac := range other.All() {
		s.Add(mac)
	}
}

// EmptySet returns a new empty set.
func EmptySet() *Set {
	return &Set{}
}

// Intersection returns a new set containing the intersection of the two given sets.
func Intersection(s1, s2 *Set) *Set {
	retv := EmptySet()
	for _, mac := range s1.All() {
		if s2.Contains(mac) {
			retv.Add(mac)
		}
	}
	for _, mac := range s2.All() {
		if s1.Contains(mac) {
			retv.Add(mac)
		}
	}
	return retv
}

// Union returns a new set containing the union of the two given sets.
func Union(s1, s2 *Set) *Set {
	retv := EmptySet()
	retv.AddAllFrom(s1)
	retv.AddAllFrom(s2)
	return retv
}
