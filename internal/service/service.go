package service

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Service(str string) string {

	for _, ch := range str {
		if ch != '.' && ch != '-' && ch != ' ' {
			return morse.ToMorse(str)

		}
	}
	return morse.ToText(str)
}
