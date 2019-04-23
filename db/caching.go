package db

import "go.mongodb.org/mongo-driver/bson"

func CacheUser(user *User) error {

	collection, err := Get("users")
	if err != nil {
		return err
	}

	dbUser, err := FindUser(user.Login, collection)
	if dbUser != nil {
		filter := bson.D{{"login", user.Login}}
		err = UpdateAll(filter, user, collection)
		if err != nil {
			return err
		}

	} else {
		err = Insert(user, collection)
		if err != nil {
			return err
		}
	}

	return nil

}
