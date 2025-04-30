package i18n

type Config struct {
	BundleDirectory string `yaml:"bundle_dir"`
}

func DefaultConfig() Config {
	return Config{
		BundleDirectory: ".i18n",
	}
}

func LoadConfig(byt []byte) (Config, error) {
	var (
		config = DefaultConfig()
		err    = yaml.Unmarshal(byt, &config)
	)
	return config, err
}
