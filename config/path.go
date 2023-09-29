package config

import (
	"fmt"
	"os/user"
	"path/filepath"
)

const (
	applicationFolder = "rename4plex"
	configFile        = "config.yml"
)

// GetConfigFilePath is a function which constructs and returns the application configuration file's absolute path.
func GetConfigFilePath() (string, error) {
	// Get the details of the user who is currently running the process.
	currentUser, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("read current user: %w", err)
	}

	// Construct the configuration file's absolute path string by using the currentUser's home directory, the application folder, and the configuration file name.
	configPath := filepath.FromSlash(fmt.Sprintf(
		"%s/.config/%s/%s",
		currentUser.HomeDir,
		applicationFolder,
		configFile,
	))

	return configPath, nil
}
