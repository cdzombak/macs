package macs

import (
	"encoding/json"
)

// Set is a (non-concurrency-safe!) set of MAC addresses which supports JSON [un]marshaling.
type Set struct {
	macs map[string]bool
}

func (s *Set) MarshalJSON() ([]byte, error) {
	keys := make([]string, 0, len(s.macs))
	for k, _ := range s.macs {
		keys = append(keys, k)
	}
	return json.Marshal(keys)
}

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

func (s *Set) Contains(mac *Addr) bool {
	s.ensureInit()
	return s.macs[mac.String()]
}

func (s *Set) Add(mac *Addr) {
	s.ensureInit()
	s.macs[mac.String()] = true
}

func (s *Set) Remove(mac *Addr) {
	s.ensureInit()
	delete(s.macs, mac.String())
}

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

func (s *Set) Len() int {
	return len(s.macs)
}

func (s *Set) AddAllFrom(other *Set) {
	s.ensureInit()
	for _, mac := range other.All() {
		s.Add(mac)
	}
}

func EmptyMACSet() *Set {
	return &Set{}
}
