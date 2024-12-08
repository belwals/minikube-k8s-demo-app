package model

type Environment struct {
	//  required env variables for running the service
	MongoClusterUrl string `json:"MONGO_CLUSTER_URL"`
	MongoUsername   string `json:"MONGO_USERNAME"`
	MongoPassword   string `json:"MONGO_PASSWORD"`
	LogLevel        string `json:"LOG_LEVEL" envDefault:"INFO"`
}
