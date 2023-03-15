package main

import "net/http"

var ErrUserInvalid = apiError{Err: "user no valid", Status: http.StatusUnauthorized}

type apiError struct {
	Err    string
	Status int
}

func (e apiError) Error() string {
	return e.Err
}
