package network

import (
	"os"

	"github.com/huandu/facebook"
)

var (
	appID      = os.Getenv("FB_ID")
	appSecret  = os.Getenv("FB_SECRET")
	appVersion = os.Getenv("FB_VERSION")

	// FbApp global facebook application
	FbApp *facebook.App
)

type (
	// FacebookGraphAPI defines graph API operations
	FacebookGraphAPI interface {
		ValidateAccessToken(accessToken string) bool
	}

	// FacebookGraphAPIClient FacebookGraphAPI implementation
	FacebookGraphAPIClient struct{}
)

// FacebookAPP init facebook application
func FacebookAPP() *facebook.App {
	facebook.Version = appVersion
	FbApp = facebook.New(appID, appSecret)

	return FbApp
}

// ValidateAccessToken validates user access token
func (fbApi *FacebookGraphAPIClient) ValidateAccessToken(accessToken string) bool {
	session := FbApp.Session(accessToken)
	err := session.Validate()

	if err == nil {
		return true
	}

	return false
}
