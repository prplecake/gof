package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"time"
)

var (
	Version    string
	Buildtime  string
	configFile string
)

type feed struct {
	URL, Template  string
	Format         string
	Visibility     string
	Sensitive      bool
	ContentWarning string
	TimeJitter     time.Duration
}

type config struct {
	Accounts    []account
	LastUpdated time.Time
	HttpConfig  httpConfig
	Meta        meta
}

type meta struct {
	Name      string
	Version   string
	Buildtime string
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
	config.Meta.Name = "gpf"
	config.Meta.Version = Version
	config.Meta.Buildtime = Buildtime
	config.HttpConfig.UserAgent = fmt.Sprintf("%s/%s",
		config.Meta.Name, config.Meta.Version)
	return config
}

func (cf *config) updateLastUpdated() {
	log.Println("updating LastUpdated key...")
	cf.LastUpdated = time.Now()  //need to think through timezones
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
