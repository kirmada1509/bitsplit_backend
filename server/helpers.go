package server

import (
	"io"
	"net/http"
)

func readBody(r *http.Request) ([]byte, error) {
	defer r.Body.Close()

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}
