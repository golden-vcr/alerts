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

	friendInputs := parseFriendInputs(s, rand.Int())
	if friendInputs != nil {
		return genreq.RequestTypeImage, &genreq.Payload{
			Image: &genreq.PayloadImage{
				Style: genreq.ImageStyleFriend,
				Inputs: genreq.ImageInputs{
					Friend: friendInputs,
				},
			},
		}, nil
	}

	return "", nil, ErrNoRequest
}
