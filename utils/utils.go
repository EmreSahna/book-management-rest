package utils

import (
	"encoding/json"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) error {
	err := json.NewDecoder(r.Body).Decode(x)
	if err != nil {
		return err
	}
	return nil
}
