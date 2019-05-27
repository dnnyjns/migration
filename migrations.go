package migration

import (
	"sort"

	"github.com/jinzhu/gorm"
)

type migrations []*Migration

func (m *migrations) Migrate(db *gorm.DB) error {
	var (
		ids       = make([]string, len(*m))
		persisted migrations
	)

	for i, migration := range *m {
		ids[i] = migration.Version
	}
	db.Model(&Migration{}).Where(ids).Order("version").Find(&persisted)

	m.sort()
	for _, migration := range *m {
		length := len(persisted)
		i := sort.Search(length, func(i int) bool { return persisted[i].Version == migration.Version })
		if i == length {
			migration.Migrate(db)
		}
	}

	return nil
}

func (m migrations) sort() {
	sort.Slice(m, func(i, j int) bool {
		return m[i].Version < m[j].Version
	})
}
