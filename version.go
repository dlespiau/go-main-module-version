package main

import (
	"errors"
	"fmt"
	"runtime/debug"
)

// version returns the version of the main module
func version() (string, error) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "", errors.New("binary has not been built with module support")
	}
	return info.Main.Version, nil
}

// Print prints the main module version on stdout
func Print() error {
	v, err := version()
	if err != nil {
		return err
	}
	fmt.Printf("Version: %#v\n", v)
	return nil
}
