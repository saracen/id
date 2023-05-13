package id

import (
	"testing"
)

func TestID(t *testing.T) {
	id, err := ID()
	if err != nil {
		t.Fatal(err)
	}

	if id == "" {
		t.Fatal("id was empty")
	}

	t.Log(id)
}
