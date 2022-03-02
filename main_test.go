package main

import (
	"fmt"
	"testing"

	"github.com/hashicorp/nomad/api"
)

func TestCreateKeyWithJobID(t *testing.T) {
	expect := "foo-service-version"
	actual := createKeyWithJobID("foo-service", "version")

	if expect != actual {
		t.Errorf("Expected %q, actual %q", expect, actual)
	}
}

func TestCreateKeyValuePairs(t *testing.T) {
	mock := make(map[string]string)
	mock["key"] = "value"

	expect := fmt.Sprintf("%s=\"%s\"\n", "key", "value")
	actual := createKeyValuePairs(mock)

	if expect != actual {
		t.Errorf("Expected %q, actual %q", expect, actual)
	}
}

func TestCreateNomadClient(t *testing.T) {
	nomadHost        := "http://localhost:4646"
	nomadCfg         := api.DefaultConfig()
	nomadCfg.Address = nomadHost
	expect, _        := api.NewClient(nomadCfg)

	actual := createNomadClient(&nomadHost)

	if expect.Address() != actual.Address() {
		t.Errorf("Expected %q, actual %q", expect.Address(), actual.Address())
	}
}

func TestCreateConsulClient(t *testing.T) {
	
}