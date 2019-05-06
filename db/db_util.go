package db

import (
	"github.com/m-lukas/github-analyser/util"
)

func getRoot() (*DatabaseRoot, error) {

	var err error

	if util.IsTesting() {
		return getTestRoot(), nil
	}

	_, err = getDefaultRoot()
	if err != nil {
		return nil, err
	}

	return dbRoot, nil
}

func getTestRoot() *DatabaseRoot {
	return TestRoot
}

func getDefaultRoot() (*DatabaseRoot, error) {

	var err error

	err = checkDbRoot()
	if err != nil {
		return nil, err
	}

	return dbRoot, nil
}

func checkDbRoot() error {

	if dbRoot == nil {
		err := Init()
		if err != nil {
			return err
		}
	}
	return nil
}
