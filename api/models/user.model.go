package models

type UsersRole string

const (
	UsersRoleAdmin UsersRole = "admin"
	UsersRoleUser  UsersRole = "user"
)

// Usuario representa um usuário
type User struct {
	Id       uint64 `json:"id,omitempty"`
	Name     string `json:"nome,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"senha,omitempty"`
	Role     UsersRole
}
