package http

import "time"

// CreateUserDTO request payload to create a new user
// swagger:parameters CreateUserDTO
type CreateUserDTO struct {
	// Email address of the user
	// example: email@email.com
	// required: true
	Email string `json:"email" example:"email@mail.com" description:"Email address of the user"`

	// User Password (8-32 characters, at least one uppercase letter, one lowercase letter, one number)
	// example: Password123!
	// required: true
	Password string `json:"password" example:"Password123!" description:"User Password (8-32 characters, at least one uppercase letter, one lowercase letter, one number)"`

	// Full name of the user
	// example: Só Mais Silva
	// required: true
	FullName string `json:"fullName" example:"Só Mais Silva" description:"Full name of the user"`
}

// UserDTO response payload to get a user
// swagger:response UserDTO
type UserDTO struct {
	// ID of the user
	// example: 123
	// required: true
	ID string `json:"id"`

	// Full name of the user
	// example: Só Mais Silva
	// required: true
	FullName string `json:"fullName"`

	// Email address of the user
	// example:email@email.com
	// required: true
	Email string `json:"email"`

	// Image of the user
	// example: http://www.user.com/image
	// required: false
	Image *string `json:"image"`
}

// LoginDTO request payload to login
// swagger:parameters LoginDTO
type LoginDTO struct {
	// Email address of the user
	//
	// example: email@mail.com
	// required: true
	Email string `json:"email" example:"email@mail.com" description:"Email address of the user"`

	// User Password (8-32 characters, at least one uppercase letter, one lowercase letter, one number)
	//
	// example: Password123!
	// required: true
	Password string `json:"password" example:"Password123!" description:"User Password (8-32 characters, at least one uppercase letter, one lowercase letter, one number)"`
}

// CreateProjectDTO request payload to create a new project
// swagger:parameters CreateProjectDTO
type CreateProjectDTO struct {
	// Name of the project
	// example: Orange Portfolio
	// required: true
	Name string `json:"name" example:"Project Name" description:"Name of the project"`

	// Description of the project
	// example: This is a project to manage portfolios
	// required: true
	Description string `json:"description" example:"Project Description" description:"Description of the project"`

	// Link to the project
	// example: http://www.project.com
	// required: false
	Link *string `json:"link" example:"http://www.project.com" description:"Link to the project"`

	// Image of the project
	// example: http://www.project.com/image
	// required: false
	Image *string `json:"image" example:"http://www.project.com/image" description:"Image of the project"`

	// Tags of the project
	// example: [1, 12]
	// required: false
	Tags []int `json:"tags" example:"[20, 30]" description:"Tag ids of the project"`
}

// ProjectDTO response payload to get all projects of a user
// swagger:response ProjectDTO
type ProjectDTO struct {
	Project ProjectInfo `json:"project"`
	User    UserInfo    `json:"user"`
	Tags    []string    `json:"tags"`
}

type ProjectInfo struct {

	// ID of the project
	// example: 123
	// required: true
	ID string `json:"id"`

	// Name of the project
	// example: Orange Portfolio
	// required: true
	Name string `json:"name"`

	// Description of the project
	// example: This is a project to manage portfolios
	// required: true
	Description string `json:"description"`

	// Link to the project
	// example: http://www.project.com
	// required: false
	Link *string `json:"link"`

	// Image of the project
	// example: http://www.project.com/image
	// required: false
	Image *string `json:"image"`

	// CreatedAt of the project
	// example: 2021-08-01T00:00:00Z
	// required: true

	CreatedAt time.Time `json:"createdAt"`

	// UpdatedAt of the project
	// example: 2021-08-01T00:00:00Z
	// required: true
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserInfo struct {

	// ID of the user
	// example: 123
	// required: true
	ID string `json:"id"`

	// Full name of the user
	// example: Só Mais Silva
	// required: true
	FullName string `json:"fullName"`

	// Email address of the user
	// example: email@email.com
	// required: true
	Email string `json:"email"`

	// Image of the user
	// example: http://www.user.com/image
	// required: false
	Image *string `json:"image"`
}

// TagDTO response payload to get all tags
// swagger:response TagDTO
type TagDTO struct {
	// ID of the tag
	// example: 123
	// required: true
	ID int `json:"id"`

	// Name of the tag
	// example: Backend
	// required: true
	Name string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
}
