package main

import (
	"testing"

	"github.com/Iteam1337/go-protobuf-wejay/version"
)

func TestVersion(t *testing.T) {
	version := version.Version
	expected := 4
	if version != expected {
		t.Errorf("version was expected to be %d, got: %d", expected, version)
	}
}
