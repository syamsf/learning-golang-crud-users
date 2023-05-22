package request

type UpdateUserInput struct {
  Name  string `json:"name" validate:"required"`
  Email string `json:"email" validate:"email"`
}
