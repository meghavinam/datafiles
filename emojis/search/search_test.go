package search

import (
	"slices"
	"testing"
)

func TestByDescription(t *testing.T) {
	tests := map[string]struct {
		Params Params
		Want   []string
	}{
		"fruit": {Params: Params{Include: []string{"face"}}, Want: []string{"smile", "wonder", "cry"}},
		"Sad":   {Params: Params{Include: []string{"face"}, Exclude: []string{"eye"}}, Want: []string{"smile", "wonder"}},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ByDescription(test.Params)
			if len(result) != len(test.Want) {
				t.Errorf("Incorrect inputs")
			}
			for _, emoji := range result {
				if !slices.Contains(test.Want, emoji.Emoji) {
					t.Errorf("Incorrect arrays")
				}
			}
		})
	}
}
