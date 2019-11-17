package main

import (
	"io/ioutil"
	"log"
	"os/user"
	"time"

	"gopkg.in/yaml.v2"
)

var (
	dir, configFile string
)

type feed struct {
	URL, Template string
	Summary       bool
}

type config struct {
	Accounts    []Account
	LastUpdated time.Time
}

// An Account holds the information required to use that account.
type Account struct {
	ClientID     string
	ClientSecret string
	AccessToken  string
	Name         string
	InstanceURL  string
	Feeds        []feed
}

func readConfig() *config {
	usr, _ := user.Current()
	dir = usr.HomeDir
	log.Println("reading config...")
	configFile = "gof.yaml"
	config := new(config)
	cf, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalln("Failed to read config: ", err)
	}
	err = yaml.Unmarshal(cf, &config)
	if err != nil {
		log.Panic(err)
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
	err = ioutil.WriteFile(configFile, cfbytes, 0644)
	if err != nil {
		log.Fatalf("Failed to save config to file. Error: %s", err.Error())
	}

	return nil
}
