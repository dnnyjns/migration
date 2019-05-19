package runner

import (
	"sort"
)

type Migraines []*Migraine

func (ms Migraines) sort() {
	sort.Slice(ms, func(i, j int) bool {
		return ms[i].Version < ms[j].Version
	})
}
