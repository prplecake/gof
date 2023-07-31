package main

import (
	"fmt"
	"testing"
)

// TestUserAgent tests User Agent
func TestUserAgent(t *testing.T) {
	conf := readConfig("gof.example.yaml")

	if conf.HttpConfig.UserAgent == "" {
		t.Error("UserAgent is blank")
	}
	expected := fmt.Sprintf("%s/%s",
		conf.Meta.Name, conf.Meta.Version)
	if conf.HttpConfig.UserAgent != expected {
		t.Error("UserAgent did not match expectation")
	}
}
