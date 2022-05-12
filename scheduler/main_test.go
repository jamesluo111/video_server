package main

import (
	"os"
	"testing"
)

func TestOS(t *testing.T) {
	os.Remove("./videos/6f6911aa-f814-45b4-923c-3efdcfcf3206")
}
