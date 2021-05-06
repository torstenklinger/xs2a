package oauth

// Tokens is the reponse from the Oauth token endpoint
type Tokens struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	TokenType    string `json:"token_type,omitempty"`
	Epires       int    `json:"expires_in,omitempty"`
}
