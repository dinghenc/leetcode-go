package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	dictionary := []string{"apple", "book", "booking", "china", "beijing", "moon"}
	root := New()
	root.BuildDictionary(dictionary)

	type field struct {
		word string
		want bool
	}
	searchFields := []field{
		{
			word: "apple",
			want: true,
		}, {
			word: "china",
			want: true,
		}, {
			word: "shanghai",
			want: false,
		},
	}
	startWithFields := []field{
		{
			word: "apple",
			want: true,
		}, {
			word: "boo",
			want: true,
		}, {
			word: "cha",
			want: false,
		},
	}

	for _, f := range searchFields {
		got := root.Search(f.word)
		assert.Equal(t, f.want, got)
	}
	for _, f := range startWithFields {
		got := root.StartWith(f.word)
		assert.Equal(t, f.want, got)
	}
}
