// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewOrder struct {
	Status       string         `json:"status"`
	Calification *int           `json:"calification"`
	Feedback     *string        `json:"feedback"`
	URL          *string        `json:"url"`
	UserID       *NewUser       `json:"user_id"`
	TechnicianID *NewTechnician `json:"technician_id"`
	TvID         *NewTelevision `json:"tv_id"`
}

type NewTechnician struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type NewTelevision struct {
	Model string `json:"model"`
	Band  string `json:"band"`
}

type NewUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Order struct {
	ID           string      `json:"id"`
	Status       string      `json:"status"`
	Calification *int        `json:"calification"`
	Feedback     *string     `json:"feedback"`
	URL          *string     `json:"url"`
	UserID       *User       `json:"user_id"`
	TechnicianID *Technician `json:"technician_id"`
	TvID         *Television `json:"tv_id"`
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type Technician struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Television struct {
	ID    string `json:"id"`
	Model string `json:"model"`
	Band  string `json:"band"`
}

type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
