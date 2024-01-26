package http

type CreateUserDTO struct {
	Email    string `json:"email" example:"email@mail.com" description:"Email address of the user"`
	Password string `json:"password" example:"Password123!" description:"User Password (8-32 characters, at least one uppercase letter, one lowercase letter, one number)"`
	FullName string `json:"fullName" example:"Só Mais Silva" description:"Full name of the user"`
}

type LoginDTO struct {
	Email    string `json:"email" example:"email@mail.com" description:"Email address of the user"`
	Password string `json:"password" example:"Password123!" description:"User Password (8-32 characters, at least one uppercase letter, one lowercase letter, one number)"`
}

type Response struct {
	Message string `json:"message"`
}
