package core

import (
	"confy/internal/globals"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/charmbracelet/log"
)


func SetConfigAndBackup(module string, newconf string) {
	
}


func BackupModule(module string) bool {
	var srcPath = filepath.Join(globals.CONFIG_HOME_PATH, module)
	var files, err = os.ReadDir(srcPath)
	if err != nil {
		log.Error(err)
		return false
	}

	var destPath = filepath.Join(globals.CONFIG_HOME_PATH, "confy", "backup", module)
	if err := os.MkdirAll(destPath, os.ModePerm); err != nil {
		log.Error(err)
		return false
	}

	for _, f := range files {
		var src = filepath.Join(srcPath, f.Name())
		var dst = filepath.Join(destPath, f.Name())

		if f.Name() == "confy" {
			log.Error("Cannot create backup of confy!")
		} else if f.IsDir() {
			if err := os.MkdirAll(dst, os.ModePerm); err != nil {
				log.Error(err)
				continue
			}
			BackupModule(filepath.Join(module, f.Name()))
		} else {
			if err := copyFile(src, dst); err != nil {
				log.Error(err)
			}
		}
	}
	return true
}


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


func copyFile(src, dst string) error {
	// simple function but i copied it off the internet anyways
        sourceFile, err := os.Open(src)
        if err != nil {
                return fmt.Errorf("failed to open source file: %w", err)
        }
        defer sourceFile.Close()

        destinationFile, err := os.Create(dst)
        if err != nil {
                return fmt.Errorf("failed to create destination file: %w", err)
        }
        defer destinationFile.Close()

        _, err = io.Copy(destinationFile, sourceFile)
        if err != nil {
                return fmt.Errorf("failed to copy file: %w", err)
        }

        err = destinationFile.Sync()
        if err != nil {
                return fmt.Errorf("failed to sync destination file: %w", err)
        }

        return nil
}
