package alerts

import "regexp"

const StaticImageIdPrayerBear = "prayer-bear"
const StaticImageIdStandBack = "stand-back"

var regexPrayerBear = regexp.MustCompile(`(?i)\bprayer[ -]?bear\b`)
var regexStandBack = regexp.MustCompile(`(?i)\bstand[ -]?back\b`)

func parseStaticImageId(s string) string {
	if regexPrayerBear.MatchString(s) {
		return StaticImageIdPrayerBear
	}
	if regexStandBack.MatchString(s) {
		return StaticImageIdStandBack
	}
	return ""
}
