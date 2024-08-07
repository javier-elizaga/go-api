package utils

import (
	"errors"
	"net/http"
	"strconv"
)

func GetFloat64Param(r *http.Request, name string) (float64, error) {
	paramStr := r.URL.Query().Get(name)
	if len(paramStr) == 0 {
		return 0, errors.New("missing mandatory param " + name)
	}
	return strconv.ParseFloat(paramStr, 64)
}
