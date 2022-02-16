package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortLinkGenerator(t *testing.T) {
	initialLink_1 := "http://google.com"
	shortLink_1 := GenerateShortLink(initialLink_1)

	initialLink_2 := "https://www.facebook.com"
	shortLink_2 := GenerateShortLink(initialLink_2)

	assert.Equal(t, shortLink_1, "Fy9z6aGA")
	assert.Equal(t, shortLink_2, "QiaaLgq2")
}
