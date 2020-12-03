package main

import (
	"testing"

	"github.com/yrisob/database_migrations/config"
)

func TestGettingConfig(t *testing.T) {
	config, err := config.GetConfig()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(config.GetConnectionString())
}


