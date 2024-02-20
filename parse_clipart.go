package alerts

import (
	"fmt"
	"regexp"
	"strings"

	genreq "github.com/golden-vcr/schemas/generation-requests"
)

var regexClipArtPrefix = regexp.MustCompile(`(?i)^!?clip(?:[- ]?art)?s?(?: of)?( an?| the)? (.+)$`)

func parseClipArtInputs(s string, randomValue int) *genreq.ImageInputsClipArt {
	// Require that the string starts with 'clip art', 'clip-art', 'clipart', etc.
	m := regexClipArtPrefix.FindStringSubmatch(s)
	if m == nil {
		return nil
	}

	// If the subject (past the article) contains a color name, normalize it to a
	// canonical color name and use it; otherwise pick a random color
	var color genreq.Color
	var subject string
	if c, remainder, err := genreq.MatchColor(m[2]); err == nil {
		color = c
		subject = remainder
	} else {
		color = genreq.Colors[randomValue%len(genreq.Colors)]
		subject = m[2]
	}

	// If there was an article in the input string, prepend it to the subject (which now
	// has no color name, e.g. 'an orange fish' will produce a color of 'orange' and a
	// subject of 'a fish')
	article := strings.TrimSpace(m[1])
	if article != "" {
		if article == "a" && beginsWithVowel(subject) {
			article = "an"
		} else if article == "an" && !beginsWithVowel(subject) {
			article = "a"
		}
		subject = fmt.Sprintf("%s %s", article, subject)
	}

	return &genreq.ImageInputsClipArt{
		Color:   color,
		Subject: subject,
	}
}

func beginsWithVowel(s string) bool {
	return len(s) > 0 && (s[0] == 'a' || s[0] == 'e' || s[0] == 'i' || s[0] == 'o' || s[0] == 'u')
}
