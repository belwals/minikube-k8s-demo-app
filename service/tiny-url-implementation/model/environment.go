package model

type Environment struct {
	//  required env variables for running the service
	MongoClusterUrl string `env:"MONGO_CLUSTER_URL,required" envDefault:"127.0.0.1:27017"`
	MongoUsername   string `env:"MONGO_USERNAME,required" envDefault:"mongouser"`
	MongoPassword   string `env:"MONGO_PASSWORD,required" envDefault:"mongopassword"`
	LogLevel        string `env:"LOG_LEVEL" envDefault:"INFO"`
	Port            int    `env:"PORT" envDefault:"8080"`
	Address         string `env:"Address" envDefault:""`
}
