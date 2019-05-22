package db

import (
	"github.com/m-lukas/github-analyser/util"
)

//getRoot returns the database root of type *DatabaseRoot
func getRoot() (*DatabaseRoot, error) {

	var err error

	//return different root in testing enviroment
	if util.IsTesting() {
		return getTestRoot(), nil
	}

	//in dev/prod return default root
	root, err := getDefaultRoot()
	if err != nil {
		return nil, err
	}

	return root, nil
}

//getTestRoot returns the testing enviroment root to be used outside the package -> avoid root initialization
func getTestRoot() *DatabaseRoot {
	return TestRoot
}

//getDefaultRoot returns a non-nil dbRoot, or error if initialization failed
func getDefaultRoot() (*DatabaseRoot, error) {

	var err error

	//retrieve dbRoot using helper function
	root, err := checkDbRoot()
	if err != nil {
		return nil, err
	}

	return root, nil
}

//checkRoot checks if dbRoot == nil and returns dbRoot if no error
func checkDbRoot() (*DatabaseRoot, error) {

	//check if root is initialized, if not initialize root and return
	if dbRoot == nil {
		err := Init()
		if err != nil {
			return nil, err
		}
	}
	return dbRoot, nil
}
