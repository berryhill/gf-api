// +build dev

package config


func MakeConfig() *Config {

	config := new(Config)

	config.MongoDBHosts = "172.17.0.1:27017"
	config.AuthDatabase = ""
	config.AuthUserName = ""
	config.AuthPassword = ""
	config.TestDatabase = "test"

	return config
}
