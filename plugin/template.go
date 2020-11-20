package plugin

var (
	tmpl = `
package main

import (
	"github.com/yadisnel/go-ms/v2/plugin"

	"{{.Path}}"
)

var Plugin = plugin.Config{
	Name: "{{.Name}}",
	Type: "{{.Type}}",
	Path: "{{.Path}}",
	NewFunc: {{.Name}}.{{.NewFunc}},
}
`
)
