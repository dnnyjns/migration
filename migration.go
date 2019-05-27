package migration

import (
	"html/template"
	"os"

	"github.com/jinzhu/gorm"
)

type performFn func(*gorm.DB) error

type Migration struct {
	DisableDDL bool      `gorm:"-"`
	Perform    performFn `gorm:"-"`
	Version    string    `gorm:"size:255;PRIMARY_KEY;NOT NULL"`
}

var (
	Dir                = "./migrations"
	internalMigrations migrations
)

func Add(m *Migration) {
	internalMigrations = append(internalMigrations, m)
}

func Create(name string) {
	m := &migrationTemplate{dir: Dir, name: name}
	m.createDir()
	tmpl, err := template.New("migraine").Parse(migrationTemplateStr)
	if err != nil {
		panic(err)
	}
	f, err := os.Create(m.file())
	defer f.Close()
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(f, *m)
	if err != nil {
		panic(err)
	}
}

func Migrate(db *gorm.DB) error {
	db.AutoMigrate(&Migration{})
	return internalMigrations.Migrate(db)
}

func (m Migration) Migrate(db *gorm.DB) {
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