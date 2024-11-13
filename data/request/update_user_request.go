package request

type UpdateUserRequest struct {
	Id    int64  `validate:"required"`
	Name  string `validate:"required,max=200,min=1" json:"name"`
	Email string `validate:"required,email" json:"email"`
}
