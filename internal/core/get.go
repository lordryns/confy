package core

import (
	"confy/internal/globals"
	"fmt"
	"os"
	"path"

	"github.com/charmbracelet/log"
)


func GetModuleDetails(config string) {
	var globalConfigPath = globals.CONFIG_HOME_PATH
	var configPath = path.Join(globalConfigPath, config) 

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
		fmt.Printf("File: %v\n", file.Name());
	}
}
