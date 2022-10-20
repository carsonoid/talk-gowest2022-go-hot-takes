package main

import (
	"encoding/json"
	"net/http"
)

// START HELPER OMIT
type HTTPHelper struct {
	w        http.ResponseWriter
	Response any
}

func NewHTTPHelper(w http.ResponseWriter) *HTTPHelper {
	return &HTTPHelper{
		w: w,
	}
}

func (h *HTTPHelper) SendJSON() {
	h.w.Header().Add("Content-Type", "application/json")
	data, _ := json.Marshal(h.Response)
	h.w.Write(data)
}

func (h *HTTPHelper) Error(err error) {
	h.w.WriteHeader(http.StatusInternalServerError)
	h.Response = map[string]any{
		"error": err.Error(),
	}
	h.SendJSON()
}

// END HELPER OMIT

// START BAD OMIT

func GetUsers(r *http.Request, w http.ResponseWriter) {
	h := NewHTTPHelper(w)

	users, err := db.GetUsers()
	if err != nil {
		h.Error(err)
		return
	}

	h.Response = users
	h.SendJSON()
}

// END BAD OMIT

// START BAD2 OMIT

func GetUsers(r *http.Request, w http.ResponseWriter) {
	h := NewHTTPHelper(w)
	defer h.SendJSON()

	users, err := db.GetUsers()
	if err != nil {
		h.Error(err)
		return
	}

	h.Response = users
}

// END BAD2 OMIT

// START FIXED OMIT
func GetUsers(r *http.Request, w http.ResponseWriter) {
	users, err := db.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-Type", "application/json")
		data, _ := json.Marshal(response)
		w.Write(data)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	data, _ := json.Marshal(users)
	w.Write(data)
}

// END FIXED OMIT

// START FIXED2 OMIT

func SendJSON(w http.ResponseWriter, code int, response any) {
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	data, _ := json.Marshal(response)
	w.Write(data)
}

func GetUsers(r *http.Request, w http.ResponseWriter) {
	h := NewHTTPHelper(w)
	defer h.SendJSON()

	users, err := db.GetUsers()
	if err != nil {
		SendJSON(w, http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
		return
	}

	SendJSON(w, http.StatusInternalServerError, http.StatusOK, users)
}

// END FIXED2 OMIT
