package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(data string) string {
	containsText := func(r rune) bool {
		return r != '.' && r != '-' && r != ' '
	}
	isText := strings.ContainsFunc(data, containsText)

	if isText {
		return morse.ToMorse(data)
	}
	return morse.ToText(data)
}
