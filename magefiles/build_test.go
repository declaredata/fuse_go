package main

import (
	"context"

	"github.com/charmbracelet/log"
	"github.com/magefile/mage/sh"
)

// Run all tests
func Test(ctx context.Context) {
	out, err := sh.Output("go", "test", "./...")
	if err != nil {
		log.Fatal(err)
	}
	log.Print(out)
	log.Info("Test success!")
}

func Build(ctx context.Context) {
	if _, err := sh.Output("go", "build", "."); err != nil {
		log.Fatal(err)
	}
	log.Info("Build success!")
}
