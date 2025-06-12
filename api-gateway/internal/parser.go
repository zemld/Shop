package internal

import (
	"log"
	"strconv"
	"strings"
)

func TryParseURLPathParamAndConvertToInt(path string, pattern string, param string) (int, error) {
	parsedValue, err := tryParseURLPathParam(path, pattern, param)
	if err != nil {
		return -1, err
	}
	parsedIntValue, err := strconv.Atoi(parsedValue)
	if err != nil {
		return -1, err
	}
	log.Printf("Parsed integer value from path: %s : %d\n", path, parsedIntValue)
	return parsedIntValue, nil
}

func tryParseURLPathParam(path string, pattern string, param string) (string, error) {
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
