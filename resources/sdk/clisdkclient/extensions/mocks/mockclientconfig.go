package mocks

import (
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/config"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/models"
)

type MockClientConfig struct {
	ProfileNameFunc    func() string
	EnvironmentFunc    func() string
	ClientIDFunc       func() string
	ClientSecretFunc   func() string
	OAuthTokenDataFunc func() string
	LogFilePathFunc    func() string
	LoggingEnabledFunc func() bool
}

var UpdatedAccessToken string

func (m *MockClientConfig) ProfileName() string {
	return m.ProfileNameFunc()
}

func (m *MockClientConfig) Environment() string {
	return m.EnvironmentFunc()
}

func (m *MockClientConfig) ClientID() string {
	return m.ClientIDFunc()
}

func (m *MockClientConfig) ClientSecret() string {
	return m.ClientSecretFunc()
}

func (m *MockClientConfig) OAuthTokenData() string {
	return m.OAuthTokenDataFunc()
}

func (m *MockClientConfig) LogFilePath() string {
	return m.LogFilePathFunc()
}

func (m *MockClientConfig) LoggingEnabled() bool {
	return m.LoggingEnabledFunc()
}

func (m *MockClientConfig) String() string {
	return fmt.Sprintf("\n-------------\nProfile Name: %s\nEnvironment: %s\nLogging Enabled: %v\nLog File Path: %s\nClient ID: %s\nClient Secret: %s\n--------------\n", m.ProfileName(), m.Environment(), m.LoggingEnabled(), m.LogFilePath(), m.ClientID(), m.ClientSecret())
}

func UpdateOAuthToken(_ config.Configuration, oauthTokenData *models.OAuthTokenData) error {
	UpdatedAccessToken = oauthTokenData.AccessToken
	return nil
}
