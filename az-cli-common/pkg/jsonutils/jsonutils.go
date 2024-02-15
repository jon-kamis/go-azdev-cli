package jsonutils

import (
	"encoding/json"
	"net/http"
)

func ReadJSON(r *http.Response, data interface{}) error {
	dec := json.NewDecoder(r.Body)

	//dec.DisallowUnknownFields()

	err := dec.Decode(data)

	if err != nil {
		return err
	}

	return nil
}
