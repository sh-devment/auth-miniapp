package main

import (
	"os"
	"reflect"
	"testing"
)

// Handlers read the active region; tests run as the zone that is deployed.
func TestMain(m *testing.M) {
	region = &regionCOM
	os.Exit(m.Run())
}

// Same check initRegion makes at startup, so a missing string fails here first.
func TestRegionsFilled(t *testing.T) {
	for name, r := range regions {
		if err := checkFilled(reflect.ValueOf(*r), "region"); err != nil {
			t.Errorf("region %s: %v", name, err)
		}
	}
}
