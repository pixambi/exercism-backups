package parsinglogfiles

import (
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	var logLevelRegex = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return logLevelRegex.MatchString(text)
}

func SplitLogLine(text string) []string {
	var logSplitRegex = regexp.MustCompile(`<[~*=-]*>`)
	return logSplitRegex.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	for _, line := range lines {
		re := regexp.MustCompile(`"([^"]*)"`)
		matches := re.FindAllStringSubmatch(line, -1)
		
		for _, match := range matches {
			if len(match) > 1 {
				quotedContent := strings.ToLower(match[1])
				if strings.Contains(quotedContent, "password") {
					count++
				}
			}
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	result := make([]string, len(lines))
	
	for i, line := range lines {
		if strings.Contains(line, "User ") {
			re := regexp.MustCompile(`User\s+(\S+)`)
			match := re.FindStringSubmatch(line)
			
			if len(match) > 1 {
				username := match[1]
				result[i] = "[USR] " + username + " " + line
			} else {
				result[i] = line
			}
		} else {
			result[i] = line
		}
	}
	
	return result
}