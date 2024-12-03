package model

type Environment struct {
	//  required env variables for running the service
	MongoDbUri string `json: "MONGODB_URI"`
}
