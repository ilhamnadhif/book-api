package config

type Config struct {
	Server   ServerConfig
	JWT      JWTConfig
	Database DatabaseConfig
	Redis    RedisConfig
}

type ServerConfig struct {
	HostPort string
}

type JWTConfig struct {
	SecretKey string `validate:"required"`
}

type DatabaseConfig struct {
	HostPort string `validate:"required"`
	Username string `validate:"required"`
	Password string
	Database string `validate:"required"`
}

type RedisConfig struct {
	HostPort string `validate:"required"`
	Password string
	DBNumber int
	TTL      int64 `validate:"required,number"`
}
