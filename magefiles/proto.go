package main

import (
	"context"
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/magefile/mage/sh"
)

// Download the proto file from the fuse_python repository.
//
// Set the FUSE_GO_PROTO_BRANCH env var to indicate the branch from which
// to download. Defaults to 'main'
func DownloadProto(ctx context.Context) {
	protoBranch, ok := os.LookupEnv("FUSE_GO_PROTO_BRANCH")
	if !ok {
		protoBranch = "main"
	}
	url := fmt.Sprintf(
		"https://raw.githubusercontent.com/declaredata/fuse_python/refs/heads/%s/proto/sds.proto",
		protoBranch,
	)
	log.Info("Downloading proto file from %s branch", protoBranch)
	if err := downloadFile(url, "fuse.proto"); err != nil {
		log.Fatal(err)
	}
	log.Info("Success!")
}

func GenProto(ctx context.Context) {
	if _, err := sh.Output("buf", "generate"); err != nil {
		log.Fatal(err)
	}
	log.Info("Generated proto files")
}
