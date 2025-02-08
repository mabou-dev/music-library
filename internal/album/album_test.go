package album

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAlbum(t *testing.T) {
	input := []string{"artist", "album"}
	want := &Album{
		Artist: "artist",
		Name:   "album",
	}

	got := New(input[0], input[1])
	assert.Equal(t, want, got)
}

func TestAlbum_GetArtist(t *testing.T) {
	input := &Album{
		Artist: "artist",
		Name:   "album",
	}
	want := input.artist

	got := input.GetArtist()
	assert.Equal(t, want, got)
}

func TestAlbum_GetName(t *testing.T) {
	input := &Album{
		Artist: "artist",
		Name:   "album",
	}
	want := input.name

	got := input.GetName()
	assert.Equal(t, want, got)
}
