package models

type Character struct {
	Name     string
	ImageUrl string
	Stats    []int
}

func (c *Character) TableName() string {
	return "character"
}
