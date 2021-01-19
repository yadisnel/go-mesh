package plugin

var (
	tmpl = `
package main

import (
	"{{.Path}}"
	"github.com/yadisnel/go-ms/v2plugins"
)

var Plugin = plugin.Plugin{
	Name: "{{.Name}}",
	Type: "{{.Type}}",
	Path: "{{.Path}}",
	NewFunc: {{.Name}}.{{.NewFunc}},
}
`
)
