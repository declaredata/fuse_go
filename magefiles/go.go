package main

import (
	"github.com/charmbracelet/log"
	"github.com/magefile/mage/sh"
)

// Run all tests.
func RunTests() {
	out, err := sh.Output("go", "test", "./...")
	if err != nil {
		log.Fatal(err)
	}

	log.Print(out)
	log.Info("Test success!")
}

// Run the Go Build.
func RunBuild() {
	if _, err := sh.Output("go", "build", "."); err != nil {
		log.Fatal(err)
	}

	log.Info("Build success!")
}
