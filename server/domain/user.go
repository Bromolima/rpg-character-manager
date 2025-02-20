package domain

type User struct {
	BaseModel
	Email    string
	Password string
	Username string
	ImageUrl string
}

type UserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,containsany=@#&%*"`
	Username string `json:"username" validate:"required,min=4,max=20"`
	ImageUrl string `json:"imageUrl"`
}
