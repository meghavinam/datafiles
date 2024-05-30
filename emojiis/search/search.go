package search

import (
	"slices"
	"strings"
)

type Params struct {
	Include []string
	Exclude []string
}


func ByDescription(params Params) (result []Emoji) {
	for _, emo := range emjis {
		if shouldExclude(emo, params.Exclude) {
			continue
		}
		for _, include := range params.Include {
			include = strings.ToLower(include)
			if strings.Contains(emo.Label, include) || slices.Contains(emo.Tags, include) {
				result = append(result, emo)
			}
		}
	}
	return
}
func shouldExclude(emo Emoji, excludes []string) bool {
	for _, exclude := range excludes {
		exclude = strings.ToLower(exclude)
		if strings.Contains(emo.Label, exclude) || slices.Contains(emo.Tags, exclude) {
			return true
		}
	}
	return false
}
