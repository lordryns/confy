package core

import (
	"confy/internal/globals"
	"os"
	"path"

	"github.com/charmbracelet/log"
)


func CreateConfyModuleIfNotExist() {
	var fullPath = path.Join(globals.CONFIG_HOME_PATH, "confy")
	var _, statErr = os.Stat(fullPath)
	if !os.IsNotExist(statErr) {
		return
	}
	if err := os.Mkdir(fullPath, os.ModePerm); err != nil {
		log.Warn("Unable to create a .config/confy path, using internal defaults")
		return
	}
}


func CheckForGlobalConfigPath() (string, error) {
	// this function exists because it might read from a config file as well to determine 
	// where the config file and only do this if the confy.json file is not available
	// confy.json not yet implemented
	var path, err = os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return path, nil
}
