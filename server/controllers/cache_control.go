package controllers

import (
	"os"
)

func HasCacheFolder() {
	cacheDir := "./cache"
	_, err := os.Stat(cacheDir)

	if os.IsNotExist(err) {
		errDir := os.Mkdir(cacheDir, 0755)
		if errDir != nil {
			logger.Infof("Error while creating folder: %s", err)
			return
		}
		logger.Infof("Success created cache folder.")
	}
}

func CleanCacheFolder() {
	cacheDir := "./cache"
	err := os.RemoveAll(cacheDir)
	if err != nil {
		logger.Infof("Error while deleting folder: %s", err)
		return
	}
	logger.Infof("Success deleted cache folder.")
}
