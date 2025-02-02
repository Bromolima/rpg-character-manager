package models

type Token struct {
}

func (t *Token) TableName() string {
	return "token"
}
