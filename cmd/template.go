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
	migraineDir = "./migraines"
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

func createTemplate(name string) {
	m := &migraineTemplate{name: name}
	tmpl, err := template.New("migraine.tmpl").ParseFiles("cmd/migraine.tmpl")
	if err != nil {
		panic(err)
	}
	f, err := os.Create(m.file())
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(f, *m)
	if err != nil {
		panic(err)
	}
}
