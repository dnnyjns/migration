package cmd

import (
	"fmt"
	"os"
	"text/template"
	"time"
)

type migraineTemplate struct {
	filename string
	name     string
	version  string
}

var (
	migraineDir         = "./migraines"
	migraineTemplateStr = `package main

import (
	r "github.com/dnnyjns/migraine/runner"
	"github.com/jinzhu/gorm"
)

// Migraine for {{.Name}}
func init() {
	r.Add(&r.Migraine{
		Version: "{{.Version}}",
		Perform: func(db *gorm.DB) error {

		},
	})
}
`
	runnerTemplateStr = []byte(`package main

import (
	r "github.com/dnnyjns/migraine/runner"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	var db *gorm.DB
	db, err := gorm.Open("postgres", "host=myhost port=myport user=gorm dbname=gorm password=mypassword")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(true)

	r.Run(db)
}
`)
)

func (m migraineTemplate) file() string {
	return fmt.Sprintf("%s/%s.go", migraineDir, m.Name())
}

func (m migraineTemplate) Name() string {
	return fmt.Sprintf("%s_%s", m.Version(), m.name)
}

func (m migraineTemplate) Version() string {
	if m.version == "" {
		m.version = time.Now().Format("20060102150405")
	}
	return m.version
}

func createDir() {
	os.MkdirAll(migraineDir, os.ModePerm)
}

func createRunner() {
	createDir()
	f, err := os.Create(fmt.Sprintf("%s/runner.go", migraineDir))
	defer f.Close()
	_, err = f.Write(runnerTemplateStr)
	if err != nil {
		panic(err)
	}
}

func createTemplate(name string) {
	createDir()
	m := &migraineTemplate{name: name}
	tmpl, err := template.New("migraine").Parse(migraineTemplateStr)
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
