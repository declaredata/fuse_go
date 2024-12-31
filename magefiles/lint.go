package main

import (
	log "github.com/charmbracelet/log"
	"github.com/magefile/mage/sh"
)

// Run linters.
//
// Requires golangci-lint to be installed. See
// https://golangci-lint.run/welcome/install/ for installation instructions.
func Lint() {
	out, err := sh.Output("golangci-lint", "run")
	if err != nil {
		log.Errorf("Linter errors:\n%s", out)
		log.Fatal(err)
	}

	log.Info("Success!")
	log.Print(out)
}

// Run the gofumpt tool against the codebase.
//
// Make sure you have gofumpt installed prior to running this. See the following
// document for installation instructions:
// https://github.com/mvdan/gofumpt?tab=readme-ov-file#gofumpt
func Fmt() {
	out, err := sh.Output("gofumpt", "-l", "-w", ".")
	if err != nil {
		log.Errorf("Format errors:\n%s", out)
		log.Fatal(err)
	}

	log.Info("Success!")
	log.Print(out)
}
