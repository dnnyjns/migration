package migration

import (
	"testing"
)

func TestSort(t *testing.T) {
	m := migrations{
		&Migration{
			Version: "2",
		},
		&Migration{
			Version: "1",
		},
	}

	m.sort()

	if m[0].Version != "1" {
		t.Errorf("m[0].Version = %s; expected 1", m[0].Version)
	}
}
