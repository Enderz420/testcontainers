package rest

import (
	"encoding/json"
	"net/http"
)


type ErrorMessage struct {
	Message any `json:"message"`
}


func ReadJSON(r *http.Request, data any) error {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(data)
	if err != nil {
		return err
	}

	return nil
}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	
	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)

	if _, err := w.Write(js); err != nil {
		return err
	}
	return nil
}

func ErrorResponse(w http.ResponseWriter, r *http.Request, message any, status int) {
	err := writeJSON(w, status, ErrorMessage{Message: message})
	if err != nil {
		w.WriteHeader(500)
	}
}

func RespondWithJSON(w http.ResponseWriter, r *http.Request, data any, status int) {
	js, err := json.Marshal(data)
	if err != nil {
		 ErrorResponse(w, r, "error", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)

	if _, err := w.Write(js); err != nil {
		ErrorResponse(w, r, "error", http.StatusInternalServerError)
	}
}