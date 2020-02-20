package migration

import (
	"sort"

	"github.com/jinzhu/gorm"
)

type migrations []*Migration

func (m *migrations) IsComplete(db *gorm.DB) bool {
	var (
		count int
		total = len(*m)
		ids   = make([]string, total)
	)

	for i, migration := range *m {
		ids[i] = migration.Version
	}
	db.Model(&Migration{}).Where(ids).Count(&count)

	return total == count
}

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
	length := len(persisted)
	for _, migration := range *m {
		version := migration.Version
		i := sort.Search(length, func(i int) bool { return persisted[i].Version >= version })
		if !(i < length && persisted[i].Version == version) {
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
