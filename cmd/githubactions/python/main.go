package main

import (
	"github.com/merico-dev/stream/internal/pkg/log"
	"github.com/merico-dev/stream/internal/pkg/plugin/githubactions/python"
)

// NAME is the name of this DevStream plugin.
const NAME = "githubactions-python"

// Plugin is the type used by DevStream core. It's a string.
type Plugin string

// Install implements the installation of some GitHub Actions workflows.
func (p Plugin) Install(options *map[string]interface{}) (bool, error) {
	return python.Install(options)
}

// Reinstall implements the installation of some GitHub Actions workflows.
func (p Plugin) Reinstall(options *map[string]interface{}) (bool, error) {
	return python.Reinstall(options)
}

// Uninstall implements the installation of some GitHub Actions workflows.
func (p Plugin) Uninstall(options *map[string]interface{}) (bool, error) {
	return python.Uninstall(options)
}

// IsHealthy implements the healthy check of GitHub Actions workflows.
func (p Plugin) IsHealthy(options *map[string]interface{}) (bool, error) {
	return python.IsHealthy(options)
}

// DevStreamPlugin is the exported variable used by the DevStream core.
var DevStreamPlugin Plugin

func main() {
	log.Infof("%T: %s is a plugin for DevStream. Use it with DevStream.\n", NAME, DevStreamPlugin)
}