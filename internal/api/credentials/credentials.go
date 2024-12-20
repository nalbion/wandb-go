package credentials

import (
	"encoding/base64"
	"net/http"
)

func NewAPIKeyCredentialProvider(apiKey string) *apiKeyCredentialProvider {
	return &apiKeyCredentialProvider{
		apiKey: apiKey,
	}
}

type apiKeyCredentialProvider struct {
	// The W&B API key
	apiKey string
}

func (c *apiKeyCredentialProvider) Apply(req *http.Request) error {
	req.Header.Set(
		"Authorization",
		"Basic "+base64.StdEncoding.EncodeToString(
			[]byte("api:"+c.apiKey)),
	)
	return nil
}
