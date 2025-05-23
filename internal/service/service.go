package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Reverse(data string) string {
	sData := strings.Split(data, "")
	var isText bool

	for s := range sData {
		if s != '.' && s != '-' {
			isText = true
			break
		}
	}

	if isText {
		return morse.ToMorse(data)
	}
	return morse.ToText(data)
}
