package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[-~*=]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`"(?i)[^"]*password[^"]*"`)
	count := 0

	for _, s := range lines {
		if re.MatchString(s) {
			count++
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\S+)`)

	for i, line := range lines {
		if re.MatchString(line) {
			matches := re.FindStringSubmatch(line)
			userName := matches[1]
			lines[i] = fmt.Sprintf("[USR] %s %s", userName, line)
		}
	}

	return lines
}
