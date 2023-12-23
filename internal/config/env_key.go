package config

// Config ...
type Config struct {
	ENV                    string `env:"ENV"`
	GOOGLE_DRIVER_AUTH     string `env:"GOOGLE_DRIVER_AUTH,required"`
	GOOGLE_DRIVER_FOLDERID string `env:"GOOGLE_DRIVER_FOLDERID,required"`
}

// IsDev ...
func (c *Config) IsDev() bool {
	return c.ENV == "develop" || c.ENV == ""
}

// IsRelease ...
func (c *Config) IsRelease() bool {
	return c.ENV == "release"
}

// IsTest ...
func (c *Config) IsTest() bool {
	return c.ENV == "test"
}
