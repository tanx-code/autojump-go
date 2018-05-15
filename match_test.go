package main

import "testing"

func TestMatchAnyway(t *testing.T) {
	input := "foo"
	paths := []string{"foo/bar", "bar/foo"}
	Assert(MatchAnyway(&input, &paths), paths, t) // expected result is same as paths
}

func TestMatchFuzzy(t *testing.T) {
	input := "br"
	paths := []string{"foo/bar"}
	Assert(MatchFuzzy(&input, &paths), paths, t) // expected result is same as paths
}

func TestMatchLast(t *testing.T) {
	input := "foo"
	paths := []string{"foo/bar", "bar/foo"}
	Assert(MatchLast(&input, &paths), []string{"bar/foo"}, t) // expected result is same as paths
}
