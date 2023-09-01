package main

import "errors"

// Define all global errors used in the application.
var (
	ErrInvalidPort = errors.New("no port parameter specified")
)
