package main

import (
	"reflect"
	"testing"

	"github.com/sahilm/fuzzy"
)

func AssertString(e string, v string, t *testing.T) {
	if !reflect.DeepEqual(e, v) {
		t.Errorf("Expected %s, got %s", e, v)
	}
}

func TestFuzzy(t *testing.T) {
	input := "api"
	paths := []string{"work/bar-api/bar.go", "toy/toy.go", "foo.go"}
	v := fuzzy.Find(input, paths)[0].Str
	AssertString("work/bar-api/bar.go", v, t)
}
