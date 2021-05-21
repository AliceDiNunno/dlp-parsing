package infra

type Config struct {
	Url string
	UserAgent string
}

func LoadConfig() Config {
	return Config{
		Url: RequireEnvString("DISNEY_API"),
		UserAgent: RequireEnvString("USER_AGENT"),
	}
}
