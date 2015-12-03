package zaif

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldGetLastPrice(t *testing.T) {
	// Load the HTML document (in real use, the type would be *goquery.Document)
	ret, err := PublicAPI.LastPrice("mona_jpy")
	assert.NoError(t, err)
	t.Log(ret, err)
}

func TestShouldGetTicker(t *testing.T) {
	// Load the HTML document (in real use, the type would be *goquery.Document)
	ret, err := PublicAPI.Ticker("mona_jpy")
	assert.NoError(t, err)
	t.Log(ret, err)
}

func TestShouldGetTrades(t *testing.T) {
	// Load the HTML document (in real use, the type would be *goquery.Document)
	ret, err := PublicAPI.Trades("mona_jpy")
	assert.NoError(t, err)
	t.Log(ret, err)
}

func TestShouldGetDepth(t *testing.T) {
	// Load the HTML document (in real use, the type would be *goquery.Document)
	ret, err := PublicAPI.Depth("mona_jpy")
	assert.NoError(t, err)
	t.Log(ret, err)
}
