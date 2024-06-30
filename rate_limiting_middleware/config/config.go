package config

type Config struct {
	ServerAddress string
}

func LoadConfig() Config {
	return Config{
		ServerAddress: "localhost:8080",
	}
}
