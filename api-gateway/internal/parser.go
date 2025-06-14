package internal

import (
	"log"
	"strings"
)

func TryParseURLPathParam(path string, pattern string, param string) (string, error) {
	index := strings.Index(pattern, param)
	if index == -1 {
		log.Printf("Parameter not found in pattern: %s : %s\n", pattern, param)
		return "", nil
	}
	parsedValue := ""
	for i := index; i < len(path) && path[i] != '/'; i++ {
		if path[i] == '{' || path[i] == '}' {
			continue
		}
		parsedValue += string(path[i])
	}
	log.Printf("Parsed value from path: %s : %s\n", path, parsedValue)
	return parsedValue, nil
}
