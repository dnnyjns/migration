package migration

import (
	"os"
	"text/template"

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

func CreateDefault(name string) {
	Create(Dir, name)
}

func Create(dir, name string) {
	m := &migrationTemplate{dir: dir, name: name}
	m.createDir()
	tmpl, err := template.New("migration").Parse(migrationTemplateStr)
	if err != nil {
		panic(err)
	}
	f, err := os.Create(m.file())
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = tmpl.Execute(f, *m)
	if err != nil {
		panic(err)
	}
}

func IsComplete(db *gorm.DB) bool {
	return internalMigrations.IsComplete(db)
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
