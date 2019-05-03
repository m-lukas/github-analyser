package db

func getRoot() (*DatabaseRoot, error) {

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
