package main

import (
	"os"
)

func commandExit(cgf *config) error {
	os.Exit(0)
	return nil
}
