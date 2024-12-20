package api

import (
	"net/http"
)

// CredentialProvider adds credentials to HTTP requests.
type CredentialProvider interface {
	// Apply sets the appropriate authorization headers or parameters on the
	// HTTP request.
	Apply(req *http.Request) error
}
