package controller

import (
	"github.com/m-lukas/github-analyser/db"
	"go.mongodb.org/mongo-driver/bson"
)

//CacheUser saves all fields of the user into the given collection
func CacheUser(user *db.User, collectionName string) error {

	mongoClient, err := db.GetMongo()
	if err != nil {
		return err
	}
	elasticClient, err := db.GetElastic()
	if err != nil {
		return err
	}
	elasticIndex := elasticClient.Config.DefaultIndex

	//check if user with login exists in collection
	mongoUser, _ := mongoClient.FindUser(user.Login, collectionName)
	//not found throws unspecific error as well as connection issues

	//get user from elastic
	elasticResultList, err := elasticClient.Search(user.Login, elasticIndex, "login")
	if err != nil {
		return err
	}

	elasticUser := &db.ElasticUser{
		Login: user.Login,
		Email: user.Email,
		Name:  user.Name,
		Bio:   user.Bio,
	}

	//elasticID for savining into mongo
	var elasticID string

	//check if user exists
	if len(elasticResultList) == 0 {

		//if user doesn't exist, insert new document
		id, err := elasticClient.Insert(elasticUser, elasticIndex)
		if err != nil {
			return err
		}

		elasticID = id

	} else {

		//COMMENTED OUT BECAUSE OF ISSUE WITH ELASTIC PACKAGE
		/*
			elasticUsers, err := db.ConvertUsers(elasticResultList)
			if err != nil {
				return err
			}

			var elasticUser *db.ElasticUser

			//take first, delete others
			for _, userObj := range elasticUsers {
				if userObj.Login == user.Login {
					elasticUser = userObj
					break
					//TODO: Delete + own function to increase readness
				}
			}


				id, err := elasticClient.Update(map[string]interface{}{"bio": user.Bio}, elasticUser.Id, elasticIndex)
				if err != nil {
					return err
				}
		*/

		//elasticID = id
	}

	//update user object
	user.ElasticID = elasticID

	//check if user was already saved into mongo
	if mongoUser != nil {

		//updata user if existing
		filter := bson.D{{"login", user.Login}}
		err = mongoClient.UpdateAll(filter, user, collectionName)
		if err != nil {
			return err
		}

	} else {

		//insert new user into collection if not existing
		err = mongoClient.Insert(user, collectionName)
		if err != nil {
			return err
		}

	}

	return nil
}

//GetUserFromCache retrieves a user from the given collection by login
func GetUserFromCache(login string, collectionName string) (*db.User, error) {

	mongoClient, err := db.GetMongo()
	if err != nil {
		return nil, err
	}

	dbUser, err := mongoClient.FindUser(login, collectionName)
	if err != nil {
		return nil, err
	}

	return dbUser, nil

}
