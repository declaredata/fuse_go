package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type BadStatusError struct {
	errCode int
}

func (w BadStatusError) Error() string {
	return fmt.Sprintf("bad status code: %d", w.errCode)
}

// downloadFile downloads a file from the specified URL and saves it to the
// specified output path.
func downloadFile(url, outputPath string) error {
	resp, err := http.Get(url) //nolint:gosec
	if err != nil {
		return fmt.Errorf("downloading file: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return BadStatusError{errCode: resp.StatusCode}
	}

	out, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("creating target proto file: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("copying downloaded file contents: %w", err)
	}

	return nil
}
