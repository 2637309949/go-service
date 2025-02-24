// Package config contains helper methods for
// client side config management (`~/.micro/config.json` file).
// It uses the `JSONValues` helper
package config

import (
	"os"

	"github.com/urfave/cli/v2"
)

var (
	// File is the filepath to the config file
	File   = "./.env"
	FileKv = map[string]string{}
)

// SetConfig sets the config file
func SetConfig(cp string) {
	File = cp
	fh, err := os.Open(File)
	if err == nil {
		defer fh.Close()
		fm, err := parse(fh)
		if err != nil {
			cli.Exit(err.Error(), 1)
		}
		FileKv = fm
	}
}

// Get a value from the .env file
func Get(key string) string {
	return FileKv[key]
}
