package image

import (
	"image/gif"
	"testing"
)

func Test(t *testing.T) {
	// Images should successfully unpack
	f, err := Root.Open("/pixel.gif")
	if err != nil {
		t.Fatal(err)
	}

	// is valid gif image?
	if _, err := gif.Decode(f); err != nil {
		t.Fatal(err)
	}
}
