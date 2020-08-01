package utils

import (
	"elephant/config"
	"elephant/log"
	"os"
)

var (
	conf   = config.GetConfig()
	logger = log.GetLogger()
)

// List out all dirs path here
var appDirs = []string{
	conf.GetString("app.path_to_log_file"),
	conf.GetString("app.path_to_storage_dir"),
}

// load dirs -
func init() {
	createDirs(appDirs)
}

// create dirs if they don't exist
func createDirs(dirs []string) {
	for _, dir := range dirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err := os.Mkdir(dir, os.ModePerm); err != nil {
				logger.Errorf("cannot create dir %v : %v", dir, err)
			}
		}
	}
}
