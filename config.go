package main

import (
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

const (
	HttpUserAgent = "gof"
)

var (
	configFile string
)

type feed struct {
	URL, Template  string
	Format         string
	Visibility     string
	Sensitive      bool
	ContentWarning string
}

type config struct {
	Accounts    []account
	LastUpdated time.Time
	HttpConfig  httpConfig
}

type account struct {
	AccessToken string
	Name        string
	InstanceURL string
	Feeds       []feed
}

type httpConfig struct {
	UserAgent string
}

func readConfig(fileName string) *config {
	log.Println("reading config...")
	configFile = fileName
	config := new(config)
	cf, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalln("Failed to read config: ", err)
	}
	err = yaml.Unmarshal(cf, &config)
	if err != nil {
		log.Panic(err)
	}
	if debug {
		log.Printf("Config:\n\n%v", config)
	}
	config.HttpConfig.UserAgent = HttpUserAgent
	return config
}

func (cf *config) updateLastUpdated() {
	log.Println("updating LastUpdated key...")
	cf.LastUpdated = time.Now()
}

func (cf *config) Save() error {
	log.Println("saving config to file...")
	cfBytes, err := yaml.Marshal(cf)
	if err != nil {
		return err
	}
	err = os.WriteFile(configFile, cfBytes, 0644)
	if err != nil {
		return err
	}
	return nil
}
