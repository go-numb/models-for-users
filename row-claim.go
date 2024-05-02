package models

type Claims struct {
	ID string `firestore:"id" json:"id,omitempty"`

	AccessToken  string `firestore:"access_token" json:"access_token,omitempty"`
	AccessSecret string `firestore:"access_secret" json:"access_secret,omitempty"`

	// Auth Request Token
	RequestToken       string `firestore:"request_token" json:"request_token,omitempty"`
	RequestTokenSecret string `firestore:"request_token_secret" json:"request_token_secret,omitempty"`

	// Referer URL
	Ref string `firestore:"ref" json:"ref,omitempty"`
	// User Class 識別子
	Classification string `firestore:"classification" json:"classification,omitempty"`
}

// GetID for interface
func (p Claims) GetID() string {
	return p.ID
}
