// +build stage

package config


func MakeConfig() *Config {

	config := new(Config)

	config.MongoDBHosts = "mongo:27017"
	config.AuthDatabase = ""
	config.AuthUserName = ""
	config.AuthPassword = ""
	config.TestDatabase = "test"

	return config
}
