package pkg

import (
	"errors"
	"strconv"
	"strings"
)

func StrictAtoi(s string) (int, error) {
	if strings.TrimLeft(s, "0") != s || strings.Contains(s, "+") {
		return 0, errors.New("strconv.Atoi: leading zeros or plus signs are not allowed")
	}
	if num, err := strconv.Atoi(s); err == nil {
		return num, nil
	} else {
		return 0, err
	}
}
