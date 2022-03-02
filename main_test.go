package main

import (
	"fmt"
	"testing"
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

}

func TestCreateConsulClient(t *testing.T) {
	
}