package main

import "testing"

func TestCreateKeyWithJobID(t *testing.T) {
	expect := "foo-service-version"
	actual := createKeyWithJobID("foo-service", "version")

	if expect != actual {
		t.Errorf("Expected %q, actual %q", expect, actual)
	}
}