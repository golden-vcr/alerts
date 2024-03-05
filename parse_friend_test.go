package alerts

import (
	"testing"

	genreq "github.com/golden-vcr/schemas/generation-requests"
	"github.com/stretchr/testify/assert"
)

func Test_parseFriendInputs(t *testing.T) {
	tests := []struct {
		s    string
		want *genreq.ImageInputsFriend
	}{
		{
			"friend blue seal",
			&genreq.ImageInputsFriend{
				Color:   genreq.ColorBlue,
				Subject: "a seal",
			},
		},
		{
			"friend who's a snail",
			&genreq.ImageInputsFriend{
				Color:   genreq.ColorPurple,
				Subject: "a snail",
			},
		},
		{
			"friends the Orange-Yellow sun setting over Canada",
			&genreq.ImageInputsFriend{
				Color:   genreq.ColorYellowOrange,
				Subject: "the sun setting over Canada",
			},
		},
		{
			"friend who is an orange fish",
			&genreq.ImageInputsFriend{
				Color:   genreq.ColorOrange,
				Subject: "a fish",
			},
		},
		{
			"friend green orangutan",
			&genreq.ImageInputsFriend{
				Color:   genreq.ColorGreen,
				Subject: "an orangutan",
			},
		},
		{
			"Cheer200 I want a friend who is blue tomato",
			&genreq.ImageInputsFriend{
				Color:   genreq.ColorBlue,
				Subject: "a tomato",
			},
		},
		{
			"chicken",
			nil,
		},
		{
			"ghost of a seal",
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got := parseFriendInputs(tt.s, 11)
			assert.Equal(t, tt.want, got)
		})
	}
}
