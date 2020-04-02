package main

import (
	"testing"

	"github.com/Iteam1337/go-protobuf-wejay/version"
)

func TestVersion(t *testing.T) {
	v := version.Version
	if v != 2 {
		t.Errorf("version was expected to be 1, got: %d", v)
	}
}
