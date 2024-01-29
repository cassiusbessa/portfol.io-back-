package http

import "time"

type CreateUserDTO struct {
	Email    string `json:"email" example:"email@mail.com" description:"Email address of the user"`
	Password string `json:"password" example:"Password123!" description:"User Password (8-32 characters, at least one uppercase letter, one lowercase letter, one number)"`
	FullName string `json:"fullName" example:"Só Mais Silva" description:"Full name of the user"`
}

type LoginDTO struct {
	Email    string `json:"email" example:"email@mail.com" description:"Email address of the user"`
	Password string `json:"password" example:"Password123!" description:"User Password (8-32 characters, at least one uppercase letter, one lowercase letter, one number)"`
}

type ProjectDTO struct {
	Project ProjectInfo `json:"project"`
	User    UserInfo    `json:"user"`
	Tags    []string    `json:"tags"`
}

type ProjectInfo struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Link        *string   `json:"link"`
	Image       *string   `json:"image"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type UserInfo struct {
	ID       string  `json:"id"`
	FullName string  `json:"fullName"`
	Email    string  `json:"email"`
	Image    *string `json:"image"`
}

type Response struct {
	Message string `json:"message"`
}
