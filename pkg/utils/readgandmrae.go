package utils

import (
	"os"
	"strings"
)

func ReadGandMrae(path string) []string{
	data, _ := os.ReadFile(path)
	files := strings.Split(string(data), "\n")
	return files
}