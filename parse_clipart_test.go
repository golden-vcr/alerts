package alerts

import (
	"testing"

	genreq "github.com/golden-vcr/schemas/generation-requests"
	"github.com/stretchr/testify/assert"
)

func Test_parseClipArtInputs(t *testing.T) {
	tests := []struct {
		s    string
		want *genreq.ImageInputsClipArt
	}{
		{
			"clip art of a blue seal",
			&genreq.ImageInputsClipArt{
				Color:   genreq.ColorBlue,
				Subject: "a seal",
			},
		},
		{
			"clipArt of a snail",
			&genreq.ImageInputsClipArt{
				Color:   genreq.ColorPurple,
				Subject: "a snail",
			},
		},
		{
			"clip-arts of the Orange-Yellow sun setting over Canada",
			&genreq.ImageInputsClipArt{
				Color:   genreq.ColorYellowOrange,
				Subject: "the sun setting over Canada",
			},
		},
		{
			"clip art of an orange fish",
			&genreq.ImageInputsClipArt{
				Color:   genreq.ColorOrange,
				Subject: "a fish",
			},
		},
		{
			"clip art of a green orangutan",
			&genreq.ImageInputsClipArt{
				Color:   genreq.ColorGreen,
				Subject: "an orangutan",
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
			got := parseClipArtInputs(tt.s, 11)
			assert.Equal(t, tt.want, got)
		})
	}
}
