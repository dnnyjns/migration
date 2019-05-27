package migration

import (
	"fmt"
	"os"
	"time"
)

type migrationTemplate struct {
	dir     string
	name    string
	version string
}

var (
	migrationTemplateStr = `package migraines

import (
	"github.com/dnnyjns/migraine"
	"github.com/jinzhu/gorm"
)

// Migraine for {{.Name}}
func init() {
	migraine.Add(&migraine.Migraine{
		Version: "{{.Version}}",
		Perform: func(db *gorm.DB) error {

		},
	})
}
`
)

func (m migrationTemplate) createDir() {
	os.MkdirAll(m.dir, os.ModePerm)
}

func (m migrationTemplate) file() string {
	return fmt.Sprintf("%s/%s.go", m.dir, m.Name())
}

func (m migrationTemplate) Name() string {
	return fmt.Sprintf("%s_%s", m.Version(), m.name)
}

func (m migrationTemplate) Version() string {
	if m.version == "" {
		m.version = time.Now().Format("20060102150405")
	}
	return m.version
}
