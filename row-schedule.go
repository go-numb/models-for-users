package models

import "time"

type TypeSchedule int

const (
	None    TypeSchedule = iota
	Yearly               // 1年ごと
	Monthly              // 1ヶ月ごと
	Weekly               // 1週間ごと
	Daily                // 1日ごと
)

func (p TypeSchedule) String() string {
	switch p {
	case Yearly:
		return "Yearly"
	case Monthly:
		return "Monthly"
	case Weekly:
		return "Weekly"
	case Daily:
		return "Daily"
	}
	return "None"
}

// Schedule is a struct for scheduling posts.
type Schedule struct {
	PostID string `csv:"-" dataframe:"post_id" firestore:"post_id,omitempty" json:"post_id,omitempty"`

	OwnerId string `csv:"-" dataframe:"owner_id" firestore:"owner_id,omitempty" json:"owner_id,omitempty"`

	IsSchedule bool `csv:"-" dataframe:"is_schedule" firestore:"is_schedule,omitempty" json:"is_schedule,omitempty"`

	// IsImmediate is a flag for immediate posting.
	// 即時投稿設定
	IsImmediate bool `csv:"-" dataframe:"is_immediate" firestore:"is_immediate,omitempty" json:"is_immediate,omitempty"`

	TypeSchedule TypeSchedule `csv:"-" dataframe:"type_schedule" firestore:"type_schedule,omitempty" json:"type_schedule,omitempty"`

	// Month is the day of the month.
	// set day
	Month time.Month `csv:"-" dataframe:"month" firestore:"month,omitempty" json:"month,omitempty"`

	// Day is the day.
	Day int `csv:"-" dataframe:"day" firestore:"day,omitempty" json:"day,omitempty"`

	// Week is the day of the week.
	// set weekday
	Week time.Weekday `csv:"-" dataframe:"week" firestore:"week,omitempty" json:"week,omitempty"`

	// Times is setting multiple times.
	Times []time.Time `csv:"-" dataframe:"times" firestore:"times,omitempty" json:"times,omitempty"`
}

// IsScheduleToday is a function to determine if the schedule is today.
// Yearly, Monthly, Weekly, Dailyは個別投稿設定扱い
func (s Schedule) IsScheduleToday(t time.Time) bool {
	switch s.TypeSchedule {
	case Yearly:
		if s.Month == 0 {
			return false
		}

		// check Yearly
		if s.Month == t.Month() && s.Day == t.Day() {
			return true
		}
	case Monthly:
		// check Monthly
		if s.Day == t.Day() {
			return true
		}
	case Weekly:
		// check Weekly
		if s.Week == t.Weekday() {
			return true
		}

	case Daily:
		// dailyの扱い
		// 投稿個別設定扱い
		if IsTime(s.Times, t) {
			return true
		}
	}

	// is daily
	return false
}

func IsTime(times []time.Time, t time.Time) bool {
	for _, v := range times {
		if v.Hour() == t.Hour() && v.Minute() == t.Minute() {
			return true
		}
	}
	return false
}
