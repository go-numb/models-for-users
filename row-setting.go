package models

import "time"

// Rule is a struct that represents the setting of a row.
type Rule struct {
	UUID string `csv:"-" dataframe:"uuid" firestore:"uuid,omitempty" json:"uuid,omitempty"`

	// IsDisenable is a flag for enable posting.
	// why: 一時的な投稿停止を行うためのフラグ。投稿停止中は投稿を行わない
	IsDisenable bool `csv:"-" dataframe:"is_disenable" firestore:"is_disenable,omitempty" json:"is_disenable,omitempty"`

	// IsImmediate is a flag for immediate posting.
	IsImmediate bool `csv:"-" dataframe:"is_immediate" firestore:"is_immediate,omitempty" json:"is_immediate,omitempty"`

	// IsScheduleAllPosts 同時刻のスケジュール予約がある場合、全ての投稿を有効にするかどうか
	// why: 複数投稿を同時刻に行う場合、全ての投稿を有効にするかどうかを設定する。年・月・週指定の同時刻投稿を行いたい。
	IsScheduleAllPosts bool `csv:"-" dataframe:"is_schedule_all_posts" firestore:"is_schedule_all_posts,omitempty" json:"is_schedule_all_posts,omitempty"`

	// 投稿期間: 投稿情報の投稿最終時間からの計画時間
	// TermHours is the default posting period.
	TermHours int `csv:"-" dataframe:"term_hours" firestore:"term_hours,omitempty" json:"term_hours,omitempty"`

	// 投稿時間群
	// Times is setting multiple times. Rule.Times > Schedule.Times
	Times []time.Time `csv:"-" dataframe:"times" firestore:"times,omitempty" json:"times,omitempty"`
}
