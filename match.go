package main

import (
	"strings"

	"github.com/sahilm/fuzzy"
)

// MatchLast usually your target dir is a suffix of an absoluted path
func MatchLast(needle *string, paths *[]string) (ret []string) {
	for _, path := range *paths {
		if strings.HasSuffix(path, *needle) {
			ret = append(ret, path)
		}
	}

	return ret
}

// MatchFuzzy is matchfuzzy
func MatchFuzzy(needle *string, paths *[]string) (ret []string) {
	matches := fuzzy.Find(*needle, *paths)
	for _, v := range matches {
		ret = append(ret, v.Str)
	}

	return ret
}

// MatchAnyway if needle is a substring of a path, return that path.
func MatchAnyway(needle *string, paths *[]string) (ret []string) {
	for _, path := range *paths {
		lowPath := strings.ToLower(path)
		if strings.Contains(lowPath, *needle) {
			ret = append(ret, path)
		}
	}

	return ret
}
