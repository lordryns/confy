package core

import (
	"fmt"
	"os"
	"path"

	"github.com/charmbracelet/log"
)

func ReplaceConfig(configName string, replPath string) error  {
	var configPath, err  = os.UserConfigDir()

	if err != nil {
		fmt.Println(err)
		log.Error("Unable to determine global config path!")
		return err
	}


	if pathExists(path.Join(configPath, configName)) {
		log.Errorf("Config %v does not exist in configuration path! create one and try again!\n", configName)
		return fmt.Errorf("Path does not exist")
	}

	if !pathExists(replPath) {
		log.Errorf("Invalid config path provided!")
		return fmt.Errorf("Path does not exist!")
	}




	return nil
}

func pathExists(path string) bool {
	var _, err = os.Stat(path);

	return !os.IsNotExist(err)
}
