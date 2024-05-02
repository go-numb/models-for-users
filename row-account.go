package models

import (
	"time"

	"github.com/google/uuid"
)

type SubscribedPlan uint8

const (
	Unsubscribed SubscribedPlan = iota
	SubscribedFree
	SubscribedBasic
	SubscribedPro
)

type Account struct {
	UUID         string `csv:"uuid" dataframe:"uuid" firestore:"uuid,omitempty" json:"uuid,omitempty"`
	ID           string `csv:"id" dataframe:"id" firestore:"id,omitempty" json:"id,omitempty"`
	Password     string `csv:"password" dataframe:"password" firestore:"password,omitempty" json:"password,omitempty"`
	SpreadID     string `csv:"spread_id" dataframe:"spread_id" firestore:"spread_id,omitempty" json:"spread_id,omitempty"`
	AccessToken  string `csv:"access_token" dataframe:"access_token" firestore:"access_token,omitempty" json:"access_token,omitempty"`
	AccessSecret string `csv:"access_secret" dataframe:"access_secret" firestore:"access_secret,omitempty" json:"access_secret,omitempty"`

	CreatedAt time.Time `csv:"created_at" dataframe:"created_at" firestore:"created_at,omitempty" json:"created_at,omitempty"`
}

// NewAccountForFree 初回会員登録用
func NewAccount(id, sheetID, accessToken, accessSecret string) *Account {
	return &Account{
		UUID:         uuid.New().String(),
		ID:           id,
		Password:     "",
		SpreadID:     sheetID,
		AccessToken:  accessToken,
		AccessSecret: accessSecret,

		CreatedAt: time.Now(),
	}
}

// GetID for interface
func (p Account) GetID() string {
	return p.ID
}
