package utils

import "strings"

// ExtractKeyValues converts OCR text into label-value pairs
func ExtractKeyValues(text string) map[string]string {
	lines := strings.Split(text, "\n")
	data := map[string]string{}

	for _, line := range lines {
		if strings.Contains(line, ":") {
			parts := strings.SplitN(line, ":", 2)
			label := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			data[label] = value
		}
	}

	return data
}
