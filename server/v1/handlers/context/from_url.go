package context

import (
	"image"
	"io"
	"net/http"
)

func GetImageFromUrl(uri string) (image.Image, string, int) {
	response, err := http.Get(uri)
	if err != nil {
		return nil, "Error when retrieving image from this uri " + uri, 400
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	img, _, err := image.Decode(response.Body)
	if err != nil {
		return nil, "Error decoding image from this uri" + uri, 500
	}

	return img, "Success", 200
}
