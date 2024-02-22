package alerts

import "regexp"

const StaticImageIdPrayerBear = "prayer-bear"

var regexPrayerBear = regexp.MustCompile(`(?i)\bprayer[ -]?bear\b`)

func parseStaticImageId(s string) string {
	if regexPrayerBear.MatchString(s) {
		return StaticImageIdPrayerBear
	}
	return ""
}
