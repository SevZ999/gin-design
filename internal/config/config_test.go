package config

import (
	"log"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig("dev")
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("%+v", config)
}
