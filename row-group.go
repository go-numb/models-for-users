package models

// Group struct: 連投投稿のグループを表す
// ParentPostId: 親投稿のID
type Group struct {
	ParentPostId string `firestore:"parentPost_id" csv:"-" json:"parent_post_id,omitempty"`
	OwnerId      string `firestore:"owner_id" csv:"-" json:"owner_id,omitempty"`

	// IsOn: 連投投稿のグループかどうか
	IsOn bool `firestore:"is_on,omitempty" csv:"-" json:"is_on,omitempty"`

	ChildPostIds []string `firestore:"post_id,omitempty" csv:"-" json:"child_post_ids,omitempty"`
}
