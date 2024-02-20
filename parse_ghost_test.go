package alerts

import (
	"testing"

	genreq "github.com/golden-vcr/schemas/generation-requests"
	"github.com/stretchr/testify/assert"
)

func Test_parseGhostInputs(t *testing.T) {
	tests := []struct {
		s    string
		want *genreq.ImageInputsGhost
	}{
		{
			"ghost of a seal",
			&genreq.ImageInputsGhost{
				Subject: "a seal",
			},
		},
		{
			"Ghost of a seal named Jerry",
			&genreq.ImageInputsGhost{
				Subject: "a seal named Jerry",
			},
		},
		{
			"ghosts of several large dogs",
			&genreq.ImageInputsGhost{
				Subject: "several large dogs",
			},
		},
		{
			"!ghost Jonathan Frakes eating a pretzel",
			&genreq.ImageInputsGhost{
				Subject: "Jonathan Frakes eating a pretzel",
			},
		},
		{
			"ghost of the concept of love",
			&genreq.ImageInputsGhost{
				Subject: "the concept of love",
			},
		},
		{
			"chicken",
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got := parseGhostInputs(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
