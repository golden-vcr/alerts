package alerts

import (
	"regexp"

	genreq "github.com/golden-vcr/schemas/generation-requests"
)

var regexGhost = regexp.MustCompile(`(?i)!?ghosts?(?: of)? (.+)$`)

func parseGhostInputs(s string) *genreq.ImageInputsGhost {
	m := regexGhost.FindStringSubmatch(s)
	if m == nil {
		return nil
	}
	subject := m[1]
	return &genreq.ImageInputsGhost{
		Subject: subject,
	}
}
