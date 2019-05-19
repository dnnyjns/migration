package runner

import (
	"github.com/jinzhu/gorm"
)

type Migraine struct {
	DisableDDL bool           `gorm:"-"`
	Perform    func(*gorm.DB) `gorm:"-"`
	Version    string         `gorm:"size:255;PRIMARY_KEY;NOT NULL"`
}

func (m Migraine) Run(db *gorm.DB) {
	var tx *gorm.DB
	if m.DisableDDL {
		tx = db
	} else {
		tx = db.Begin()
	}
	defer tx.Commit()
	m.Perform(tx)

	if err := tx.Create(&m).Error; err != nil {
		tx.Rollback()
	}
}
