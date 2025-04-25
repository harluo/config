package internal

type Tag struct {
	Default string
}

func NewTag() *Tag {
	return &Tag{
		Default: "default",
	}
}
