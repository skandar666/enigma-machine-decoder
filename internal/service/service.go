package service

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Service(str []byte) string {

	for _, ch := range str {
		if ch != '.' && ch != '-' && ch != ' ' {
			return morse.ToMorse(string(str))

		}
	}
	return morse.ToText(string(str))
}
