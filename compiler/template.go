package compiler

const (
	programTemplate = `
package main

import (
	"tora/component"
	{{range .Imports}}"{{.}}"
	{{end}}
)

var (
	comps    = &component.Components{}
	services = &component.Services{}
)

func Setup() {
	// Registrations of components
	{{range .Components}}comps.Register({{.}})
	{{end}}
	// Setup services
	services.Setup(comps)
}

func List() {
	services.List()
}

func Handle(cmd string, data []byte) error {
	return services.Handle(cmd, data)
}
	`
)
