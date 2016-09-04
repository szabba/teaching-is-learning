// Use
//
//     go run mkdir_today.go
//
// to create a directory to publish a draft in, when you're ready.

package main

import (
	"os"
	"path/filepath"
	"time"
)

func main() {
	today := time.Now().UTC()
	os.MkdirAll(
		filepath.Join(
			today.Format("2006"),
			today.Format("01"),
			today.Format("02"),
		),
		0755,
	)
}
