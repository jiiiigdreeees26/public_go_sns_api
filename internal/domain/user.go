package domain

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email,omitempty"`
	Picture  string `json:"picture,omitempty"`
	Auth0Sub string `json:"auth0_sub,omitempty"`
}

type PublicUser struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
