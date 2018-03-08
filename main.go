package main

import (
	"log"
)

// Version contains the package version
var Version = "0.0.0-default"

// Main entry point for the app.
func main() {
	log.Printf("Version %q", Version)
}
