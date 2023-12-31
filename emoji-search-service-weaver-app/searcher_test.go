package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/ServiceWeaver/weaver/weavertest"
	"github.com/google/go-cmp/cmp"
)

func TestSearch(t *testing.T) {
	type test struct {
		query string
		want  []string
	}

	var testCases = []test{
		{"pig", []string{"🐖", "🐗", "🐷", "🐽"}},
		{"Pig", []string{"🐖", "🐗", "🐷", "🐽"}},
		{"black cat", []string{"🐈\u200d⬛"}},
		{"foo bar baz", nil},
	}

	for _, test := range testCases {
		for _, runner := range weavertest.AllRunners() {
			runner.Name = fmt.Sprintf("%s%q", runner.Name, test.query)
			runner.Test(t, func(t *testing.T, searcher Searcher) {
				got, err := searcher.Search(context.Background(), test.query)
				if err != nil {
					t.Fatalf("Search: %v", err)
				}
				if diff := cmp.Diff(test.want, got); diff != "" {
					t.Fatalf("Search (-want,+got):\n%s", diff)
				}
			})
		}
	}
}
