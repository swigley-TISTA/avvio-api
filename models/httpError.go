package models

type HTTPError struct {
	Error     string `json:"error"`
	Description string `json:"description"`
}
