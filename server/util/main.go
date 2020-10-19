package util

import (
	"strings"
)

func Trimspace(s *string) *string {
	if s != nil {
		*s = strings.TrimSpace(*s)
	}
	return s
}
