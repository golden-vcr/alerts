package alerts

import (
	"testing"

	eonscreen "github.com/golden-vcr/schemas/onscreen-events"
	"github.com/stretchr/testify/assert"
)

func Test_ParseStaticAlert(t *testing.T) {
	tests := []struct {
		s    string
		want *eonscreen.ImageDetailsStatic
	}{
		{
			"may Prayer bear bless us, every one",
			&eonscreen.ImageDetailsStatic{
				ImageId: "prayer-bear",
				Message: "may Prayer bear bless us, every one",
			},
		},
		{
			"prayerbear is great",
			&eonscreen.ImageDetailsStatic{
				ImageId: "prayer-bear",
				Message: "prayerbear is great",
			},
		},
		{
			"thanks be to prayer-BEAR",
			&eonscreen.ImageDetailsStatic{
				ImageId: "prayer-bear",
				Message: "thanks be to prayer-BEAR",
			},
		},
		{
			"I think you need to STAND BACK",
			&eonscreen.ImageDetailsStatic{
				ImageId: "stand-back",
				Message: "I think you need to STAND BACK",
			},
		},
		{
			"none of these words are alert keywords",
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got, err := ParseStaticAlert(tt.s)
			if tt.want != nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			} else {
				assert.ErrorIs(t, err, ErrNoStaticAlert)
				assert.Nil(t, got)
			}
		})
	}
}
