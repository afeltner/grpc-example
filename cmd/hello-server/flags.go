package main

import (
	"flag"

	"github.com/hashicorp/go-multierror"
)

//nolint:gochecknoglobals
var (
	port = flag.Int("port", 0, "The server listen port")
)

// verifyProgramArguments will verify that all required program arguments have been supplied.
// It is set up to return multiple errors, although currently there are only 1 program argument.
func verifyProgramArguments() error {
	var errs *multierror.Error
	if *port == 0 {
		errs = multierror.Append(errs, ErrInvalidPort)
	}
	return errs.ErrorOrNil()
}
