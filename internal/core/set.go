package core

import (
	"confy/internal/globals"
	"fmt"
	"os"
	"path"

	"github.com/charmbracelet/log"
)

func ReplaceModule(configName string, replPath string) error  {
	if pathExists(path.Join(globals.CONFIG_HOME_PATH, configName)) {
		log.Errorf("Config %v does not exist in configuration path! create one and try again!\n", configName)
		return fmt.Errorf("Path does not exist")
	}

	if !pathExists(replPath) {
		log.Errorf("Invalid module path  provided!")
		return fmt.Errorf("Path does not exist!")
	}




	return nil
}

func pathExists(path string) bool {
	var _, err = os.Stat(path);

	return !os.IsNotExist(err)
}
