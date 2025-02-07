package album

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAlbum(t *testing.T) {
	input := []string{"artist", "album"}
	want := &Album{
		artist: "artist",
		name:   "album",
		genres: []string(nil),
	}

	got := New(input[0], input[1])
	assert.Equal(t, want, got)
}

func TestAlbum_GetArtist(t *testing.T) {
	input := &Album{
		artist: "artist",
		name:   "album",
		genres: []string(nil),
	}
	want := input.artist

	got := input.GetArtist()
	assert.Equal(t, want, got)
}

func TestAlbum_GetName(t *testing.T) {
	input := &Album{
		artist: "artist",
		name:   "album",
		genres: []string(nil),
	}
	want := input.name

	got := input.GetName()
	assert.Equal(t, want, got)
}
