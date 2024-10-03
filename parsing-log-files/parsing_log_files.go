package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {

	regExp := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\].+`)

	return regExp.MatchString(text)
}

func SplitLogLine(text string) []string {
	regExp := regexp.MustCompile(`<[~\*=-]*>`)

	return regExp.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	regExp := regexp.MustCompile(`(?i)["'][^"']*password[^"']*["']`)

	count := 0
	for _, line := range lines {
		if !regExp.MatchString(line) {
			continue
		}

		count++
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	regExp := regexp.MustCompile(`end-of-line[0-9]+`)
	if !regExp.MatchString(text) {
		return text
	}

	return regExp.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {

	regExp := regexp.MustCompile(`User\s+([^\s]+)`)

	for idx, value := range lines {

		result := regExp.FindStringSubmatch(value)

		if len(result) < 2 {
			continue
		}

		lines[idx] = "[USR] " + result[1] + " " + value
	}

	return lines
}
