package model

type Task struct {
	ID        string `json:"id" bson:"_id"`
	Title     string `json:"title" bson:"title"`
	Slug      string `json:"slug" bson:"slug,omitempty"`
	Note      string `json:"note" bson:"note"`
	Completed bool   `json:"completed" bson:"completed"`
	CreatedAt int64  `json:"created_at" bson:"created_at"`
	UpdatedAt int64  `json:"updated_at" bson:"updated_at"`
}

func (t *Task) IsBaseModel() {}

func (t *Task) GetID() string {
	return t.ID
}

func (t *Task) GetTitle() string {
	return t.Title
}

func (t *Task) GetSlug() *string {
	return &t.Slug
}

func (t *Task) GetCreatedAt() *int {
	val := int(t.CreatedAt)
	return &val
}

func (t *Task) GetUpdatedAt() *int {
	val := int(t.UpdatedAt)
	return &val
}
