package runner

import (
	"github.com/jinzhu/gorm"
)

type performFn func(*gorm.DB) error

type Migraine struct {
	DisableDDL bool      `gorm:"-"`
	Perform    performFn `gorm:"-"`
	Version    string    `gorm:"size:255;PRIMARY_KEY;NOT NULL"`
}

var (
	internalMigraines migraines
)

func (m Migraine) Run(db *gorm.DB) {
	var tx *gorm.DB
	if m.DisableDDL {
		tx = db
	} else {
		tx = db.Begin()
		defer tx.Commit()
	}

	if err := m.Perform(tx); err != nil {
		tx.Rollback()
		panic(err)
	}

	if err := tx.Create(&m).Error; err != nil {
		tx.Rollback()
		panic(err)
	}
}

func Add(m *Migraine) {
	internalMigraines = append(internalMigraines, m)
}

func Run(db *gorm.DB) error {
	db.AutoMigrate(&Migraine{})
	return internalMigraines.Run(db)
}
