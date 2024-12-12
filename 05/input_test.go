package main

import (
	"testing"
)

func TestInput(t *testing.T) {
	v := ContentToValidator(ex1)
	if len(v.S) != 6 {
		t.Error("wrong number of rules")
	}
	if len(v.M) != 6 {
		t.Error("wrong number of manuals")
	}
}
