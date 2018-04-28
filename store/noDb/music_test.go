package noDb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchByIndexReturnsValue(t *testing.T) {
	actualResult := NewMusicStore().SearchById("5")
	assert.NotEqual(t, actualResult, nil)
	assert.Equal(t, actualResult.Album, "Rock Or Burst")
	assert.Equal(t, actualResult.Title, "Got Some Rock & Roll Thunder")
	assert.Equal(t, actualResult.Id, "5")
	assert.Equal(t, actualResult.Artist, "AC/DC")
	assert.Equal(t, actualResult.Genre, "Rock")
}

func TestSearchByIndexReturnsNoValue(t *testing.T) {
	actualResult := NewMusicStore().SearchById("-1")
	assert.Nil(t, actualResult)
}
