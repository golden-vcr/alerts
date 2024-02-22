package alerts

import (
	"errors"
	"math/rand"

	genreq "github.com/golden-vcr/schemas/generation-requests"
	eonscreen "github.com/golden-vcr/schemas/onscreen-events"
)

var ErrNoGenerationRequest = errors.New("not a generation request")
var ErrNoStaticAlert = errors.New("no static alerts")

func ParseGenerationRequest(s string) (genreq.RequestType, *genreq.Payload, error) {
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

	return "", nil, ErrNoGenerationRequest
}

func ParseStaticAlert(s string) (*eonscreen.ImageDetailsStatic, error) {
	imageId := parseStaticImageId(s)
	if imageId != "" {
		return &eonscreen.ImageDetailsStatic{
			ImageId: imageId,
			Message: s,
		}, nil
	}

	return nil, ErrNoStaticAlert
}
