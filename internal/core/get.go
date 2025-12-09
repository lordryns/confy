package core

import (
	"confy/internal/globals"
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
)

func DoesModuleExist(module string) bool {
	var m = filepath.Join(globals.CONFIG_HOME_PATH, module)
	var _, err = os.Stat(m) 
	return !os.IsNotExist(err)
}

func GetModuleDetails(config string) {
	var globalConfigPath = globals.CONFIG_HOME_PATH
	var configPath = filepath.Join(globalConfigPath, config) 

	if !pathExists(configPath) {
		log.Errorf("Unable to find config '%v'", config)
		fmt.Println("Hint: Are you sure this config exists? Check if you're reading the correct .config path")
		return 
	}

	var files, err = os.ReadDir(configPath)
	if err != nil {
		log.Error(err)
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Printf("Module: %v\n", file.Name());
		} else {
			fmt.Printf("Config: %v\n", file.Name());
		}
	}
}
