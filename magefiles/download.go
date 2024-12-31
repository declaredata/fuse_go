package main

import (
	"context"
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
func downloadFile(ctx context.Context, url, outputPath string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("creating GET %s request: %w", url, err)
	}

	resp, err := http.DefaultClient.Do(req)
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
