package queue

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Enqueue[T any](w http.ResponseWriter, r *http.Request, status int, v T) error {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}

	return nil
}

func Dequeue[T any](r *http.Request) (T, error) {
	var v T

	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}

	return v, nil
}
