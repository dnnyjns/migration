package runner

import (
	"sort"

	"github.com/jinzhu/gorm"
)

type migraines []*Migraine

func (m *migraines) Run(db *gorm.DB) error {
	var (
		ids       = make([]string, len(*m))
		persisted migraines
	)

	for i, migraine := range *m {
		ids[i] = migraine.Version
	}
	db.Model(&Migraine{}).Where(ids).Order("version").Find(&persisted)

	m.sort()
	for _, migraine := range *m {
		length := len(persisted)
		i := sort.Search(length, func(i int) bool { return persisted[i].Version == migraine.Version })
		if i == length {
			migraine.Run(db)
		}
	}

	return nil
}

func (m migraines) sort() {
	sort.Slice(m, func(i, j int) bool {
		return m[i].Version < m[j].Version
	})
}
