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
	migrationTemplateStr = `package migrations

import (
	"github.com/dnnyjns/migration"
	"github.com/jinzhu/gorm"
)

// Migration for {{.Name}}
func init() {
	migration.Add(&migration.Migration{
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
