package model

type Environment struct {
	//  required env variables for running the service
	MongoDbUri string `json:"MONGO_DB_URI"`
	LogLevel   string `json:"LOG_LEVEL" envDefault:"INFO"`
}
