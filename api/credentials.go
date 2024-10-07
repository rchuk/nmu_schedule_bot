package api

type Credentials interface {
	Username() string
	Password() string
	ApiVersion() string
	AccessToken() *string
	SetAccessToken(accessToken string)
}

type BasicCredentials struct {
	username    string
	password    string
	apiVersion  string
	accessToken *string
}

func NewCredentials(username string, password string, apiVersion string) BasicCredentials {
	return BasicCredentials{
		username:   username,
		password:   password,
		apiVersion: apiVersion,
	}
}

func (credentials *BasicCredentials) Username() string {
	return credentials.username
}

func (credentials *BasicCredentials) Password() string {
	return credentials.password
}

func (credentials *BasicCredentials) ApiVersion() string {
	return credentials.apiVersion
}

func (credentials *BasicCredentials) AccessToken() *string {
	return credentials.accessToken
}

func (credentials *BasicCredentials) SetAccessToken(accessToken string) {
	credentials.accessToken = &accessToken
}
