package dto
type LoginUserRequestDTO struct{
	Email string  	`json:"email"  validate:"required,email"`
	Password string		`json:"password"   validate:"required,min=4" `
}
type GetUserByIdDTO struct{
	Id int64  `json:"id"  validate:"required"`
}

type CreateUserDTO struct{
	Username string   `json:"username"  validate:"required"`
	Email string   	 `json:"email"  validate:"required"`
	Password string			 `json:"password"  validate:"required,min=8"`
}