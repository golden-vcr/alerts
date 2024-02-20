package alerts

import (
	"errors"
	"math/rand"

	genreq "github.com/golden-vcr/schemas/generation-requests"
)

var ErrNoRequest = errors.New("not a generation request")

func ParseRequest(s string) (genreq.RequestType, *genreq.Payload, error) {
	ghostInputs := parseGhostInputs(s)
	if ghostInputs != nil {
		return genreq.RequestTypeImage, &genreq.Payload{
			Image: &genreq.PayloadImage{
				Style: genreq.ImageStyleGhost,
				Inputs: genreq.ImageInputs{
					Ghost: ghostInputs,
				},
			},
		}, nil
	}

	clipArtInputs := parseClipArtInputs(s, rand.Int())
	if clipArtInputs != nil {
		return genreq.RequestTypeImage, &genreq.Payload{
			Image: &genreq.PayloadImage{
				Style: genreq.ImageStyleClipArt,
				Inputs: genreq.ImageInputs{
					ClipArt: clipArtInputs,
				},
			},
		}, nil
	}

	return "", nil, ErrNoRequest
}
