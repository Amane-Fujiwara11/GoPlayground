package validation

import (
	"encoding/json"
	"net/http"
)

func ParseAndValidateJSON(r *http.Request, dst interface{}, validateFunc func(interface{}) error) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(dst); err != nil {
		return err
	}
	if validateFunc != nil {
		return validateFunc(dst)
	}
	return nil
}
