package models

import (
	"time"
)

// Subscribe is used to manage the subscription status of the account
// アカウントの購読状況を管理するために使用されます
type Subscribe struct {
	// ID is AccountID
	ID string `json:"id,omitempty"`
	// Plan is SubscribedPlan
	Plan SubscribedPlan `json:"plan,omitempty"`

	ManagedGUI Managed `json:"managed_gui,omitempty"`
	ManagedAPI Managed `json:"managed_api,omitempty"`
}

type Managed struct {
	Limit Count `json:"limit,omitempty"`
	Used  Count `json:"used,omitempty"`

	LastUsedAt time.Time `json:"last_used_at,omitempty"`
}

type Count struct {
	Monthly uint16 `json:"monthly,omitempty"`
	Daily   uint16 `json:"daily,omitempty"`
	Hourly  uint16 `json:"hourly,omitempty"`
}

// NewSubscribe is constructor
// 初期設定値を返し、PlanlevelはUnsubscribedで使用回数0を返す
func NewSubscribe() *Subscribe {
	return &Subscribe{
		Plan: Unsubscribed,
		ManagedGUI: Managed{
			Limit: Count{
				Monthly: 0,
				Daily:   0,
				Hourly:  0,
			},
			Used: Count{
				Monthly: 0,
				Daily:   0,
				Hourly:  0,
			},
		},
		ManagedAPI: Managed{
			Limit: Count{
				Monthly: 0,
				Daily:   0,
				Hourly:  0,
			},
			Used: Count{
				Monthly: 0,
				Daily:   0,
				Hourly:  0,
			},
		},
	}
}

// Set is used to set the usage limit according to the subscribed plan
// 購読プランに応じて使用制限を設定するために使用されます
func (s *Subscribe) Set(level SubscribedPlan) *Subscribe {
	s.Plan = level

	switch level {
	case SubscribedFree:
		s.ManagedGUI.Limit.Monthly = 0
		s.ManagedGUI.Limit.Daily = 0
		s.ManagedGUI.Limit.Hourly = 0

		s.ManagedAPI.Limit.Monthly = 50
		s.ManagedAPI.Limit.Daily = 3
		s.ManagedAPI.Limit.Hourly = 1

	case SubscribedBasic:
		s.ManagedGUI.Limit.Monthly = 30
		s.ManagedGUI.Limit.Daily = 1
		s.ManagedGUI.Limit.Hourly = 1

		s.ManagedAPI.Limit.Monthly = 10000
		s.ManagedAPI.Limit.Daily = 1000
		s.ManagedAPI.Limit.Hourly = 100

	case SubscribedPro:
		s.ManagedAPI.Limit.Monthly = 1000
		s.ManagedAPI.Limit.Daily = 100
		s.ManagedAPI.Limit.Hourly = 10

		s.ManagedAPI.Limit.Monthly = 10000
		s.ManagedAPI.Limit.Daily = 1000
		s.ManagedAPI.Limit.Hourly = 100

	}

	return s
}

// Increment is used to increment the usage count
// 使用回数をインクリメントするために使用されます
// 月、日、時の比較を行い、使用回数をリセットし、最終使用日時を更新します
func (s *Subscribe) Increment(isAPI bool) {
	now := time.Now()
	if isAPI {
		// 最終使用時間が0でない場合(as 使用している場合)、月、日、時の使用回数をリセット
		if !s.ManagedAPI.LastUsedAt.IsZero() {
			// 月が変わった場合、月の使用回数をリセットします。
			// 日が変わった場合、または月が変わった場合、日の使用回数をリセットします。
			// 時間が変わった場合、日が変わった場合、または月が変わった場合、時間の使用回数をリセットします。
			// これにより、月が変わった場合でも、日と時間の使用回数が確実にリセットされます。
			if now.Month() != s.ManagedAPI.LastUsedAt.Month() {
				s.ManagedAPI.Used.Monthly = 0
			}
			if now.Day() != s.ManagedAPI.LastUsedAt.Day() ||
				now.Month() != s.ManagedAPI.LastUsedAt.Month() {
				s.ManagedAPI.Used.Daily = 0
			}
			if now.Hour() != s.ManagedAPI.LastUsedAt.Hour() ||
				now.Day() != s.ManagedAPI.LastUsedAt.Day() ||
				now.Month() != s.ManagedAPI.LastUsedAt.Month() {
				s.ManagedAPI.Used.Hourly = 0
			}
		}

		// 使用回数をインクリメントし、最終使用日時を更新
		s.ManagedAPI.Used.Monthly++
		s.ManagedAPI.Used.Daily++
		s.ManagedAPI.Used.Hourly++
		s.ManagedAPI.LastUsedAt = now

	} else { // GUI
		// 最終使用時間が0でない場合(as 使用している場合)、月、日、時の使用回数をリセット
		if !s.ManagedGUI.LastUsedAt.IsZero() {
			// 月が変わった場合、月の使用回数をリセットします。
			// 日が変わった場合、または月が変わった場合、日の使用回数をリセットします。
			// 時間が変わった場合、日が変わった場合、または月が変わった場合、時間の使用回数をリセットします。
			// これにより、月が変わった場合でも、日と時間の使用回数が確実にリセットされます。
			if now.Month() != s.ManagedGUI.LastUsedAt.Month() {
				s.ManagedGUI.Used.Monthly = 0
			}
			if now.Day() != s.ManagedGUI.LastUsedAt.Day() ||
				now.Month() != s.ManagedGUI.LastUsedAt.Month() {
				s.ManagedGUI.Used.Daily = 0
			}
			if now.Hour() != s.ManagedGUI.LastUsedAt.Hour() ||
				now.Day() != s.ManagedGUI.LastUsedAt.Day() ||
				now.Month() != s.ManagedGUI.LastUsedAt.Month() {
				s.ManagedGUI.Used.Hourly = 0
			}
		}

		// 使用回数をインクリメントし、最終使用日時を更新
		s.ManagedGUI.Used.Monthly++
		s.ManagedGUI.Used.Daily++
		s.ManagedGUI.Used.Hourly++
		s.ManagedGUI.LastUsedAt = now

	}
}

// IsLimit is used to check if the usage limit has been reached
// 使用制限に達したかどうかを確認するために使用されます
func (s *Subscribe) IsLimit(isAPI bool) bool {
	if isAPI {
		if s.ManagedAPI.Used.Monthly > s.ManagedAPI.Limit.Monthly {
			return true
		} else if s.ManagedAPI.Used.Daily > s.ManagedAPI.Limit.Daily {
			return true
		} else if s.ManagedAPI.Used.Hourly > s.ManagedAPI.Limit.Hourly {
			return true
		}

	} else {
		if s.ManagedGUI.Used.Monthly > s.ManagedGUI.Limit.Monthly {
			return true
		} else if s.ManagedGUI.Used.Daily > s.ManagedGUI.Limit.Daily {
			return true
		} else if s.ManagedGUI.Used.Hourly > s.ManagedGUI.Limit.Hourly {
			return true
		}

	}

	return false
}
