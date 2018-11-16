package util

// Configuration holds the root of the configuration file
type Configuration struct {
	DB     DbConfig
	Server ServerConfig
	Crypto CryptoConfig
}
