package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"address" env-required:"true"`
}

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"` //struct tags
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		flags := flag.String("config", "", "path to the config file") //location,default value, desc
		flag.Parse()

		configPath = *flags //dereferncing the flag pointer
		if configPath == "" {
			log.Fatal("Config path is not set")
		}
	}
	//if file is available at the given path

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file doesnot exist: %s", configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("Cannot read config file: %s", err.Error())
	}

	return &cfg
}
