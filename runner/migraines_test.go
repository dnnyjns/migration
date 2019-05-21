package runner

import (
	"testing"
)

func TestSort(t *testing.T) {
	m := migraines{
		&Migraine{
			Version: "2",
		},
		&Migraine{
			Version: "1",
		},
	}

	m.sort()

	if m[0].Version != "1" {
		t.Errorf("m[0].Version = %s; expected 1", m[0].Version)
	}
}
