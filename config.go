package main

import (
	"log"
	"os"
	"os/user"
	"time"

	"gopkg.in/yaml.v2"
)

var (
	dir, configFile string
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
}

type account struct {
	AccessToken string
	Name        string
	InstanceURL string
	Feeds       []feed
}

func readConfig(fileName string) *config {
	usr, _ := user.Current()
	dir = usr.HomeDir
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
	return config
}

func (cf *config) updateLastUpdated() {
	log.Println("updating lastupdated key...")
	cf.LastUpdated = time.Now()
}

func (cf *config) Save() error {
	log.Println("saving config to file...")
	cfbytes, err := yaml.Marshal(cf)
	if err != nil {
		log.Fatalln("Failed to marshal config: ", err.Error())
	}
	err = os.WriteFile(configFile, cfbytes, 0644)
	if err != nil {
		log.Fatalf("Failed to save config to file. Error: %s", err.Error())
	}

	return nil
}
