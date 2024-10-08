package compiler

const (
	programTemplate = `
// -----------------------------------------------------------
//  This program is auto generated by tora compiler. Please
//  do not edit this file.
// -----------------------------------------------------------

package main

import (
	"github.com/nanbunongjaya/tora/component"
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

func List() map[string][]string {
	return services.List()
}

func Handle(cmd string, data []byte) error {
	return services.Handle(cmd, data)
}
	`
)
