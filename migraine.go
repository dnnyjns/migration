package migraine

import (
	r "github.com/dnnyjns/migraine/runner"
	"github.com/jinzhu/gorm"
)

type Migraine r.Migraine

var (
	migraines = make(r.Migraines, 0)
)

func Add(m *Migraine) {
	migraines = append(migraines, m.toRunner())
}

func Runner(db *gorm.DB) *r.Runner {
	db.AutoMigrate(&r.Migraine{})
	return &r.Runner{db, migraines}
}

func (m Migraine) toRunner() *r.Migraine {
	return &r.Migraine{
		m.DisableDDL,
		m.Perform,
		m.Version,
	}
}
