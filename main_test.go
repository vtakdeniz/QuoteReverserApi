package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseQuotesToMap(t *testing.T) {
	type args struct {
		quotes *[]Quote
	}
	tests := []struct {
		name string
		args args
		want AuthorToQuoteMap
	}{
		{
			name: "Matches authors to quotes",
			args: args{
				quotes: &[]Quote{
					{
						Text:   "foo",
						Author: "Thomas Edison",
					},
					{
						Text:   "bar",
						Author: "Lao Tzu",
					},
					{
						Text:   "baz",
						Author: "Thomas Edison",
					},
				},
			},
			want: map[string][]string{
				"Thomas Edison": {
					"foo", "baz",
				},
				"Lao Tzu": {
					"bar",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseQuotesToMap(tt.args.quotes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseQuotesToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseQuotes(t *testing.T) {
	type args struct {
		authorToQuoteMap AuthorToQuoteMap
	}
	tests := []struct {
		name    string
		args    args
		want    AuthorToQuoteMap
		NotWant AuthorToQuoteMap
	}{
		{
			name: "Reverses each authors quote",
			args: args{
				authorToQuoteMap: map[string][]string{
					"Thomas Edison": {
						"foo", "baz",
					},
					"Lao Tzu": {
						"bar",
					},
				},
			},
			want: AuthorToQuoteMap{
				"Thomas Edison": {
					"oof", "zab",
				},
				"Lao Tzu": {
					"rab",
				},
			},
			NotWant: AuthorToQuoteMap{
				"Thomas Edison": {
					"foo", "baz",
				},
				"Lao Tzu": {
					"bar",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseQuotes(tt.args.authorToQuoteMap)
			assert.Equal(t, tt.args.authorToQuoteMap, tt.want, tt.name)
			assert.NotEqual(t, tt.args.authorToQuoteMap, tt.NotWant, tt.name)
		})
	}
}

func Test_generateQuoteJSON(t *testing.T) {
	type args struct {
		m AuthorToQuoteMap
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Tests if given map turned into result json format",
			args: args{
				m: map[string][]string{
					"Thomas Edison": {
						"foo", "baz",
					},
					"Lao Tzu": {
						"bar",
					},
				},
			},
			want: `{"author":"Thomas Edison","quotes":["foo","baz"]}{"author":"Lao Tzu","quotes":["bar"]}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateQuoteJSON(tt.args.m); got != tt.want {
				t.Errorf("generateQuoteJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
