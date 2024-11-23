package domain

type Role int32

const (
	RoleAdmin Role = iota
	RoleUser
)

type User struct {
	Id    string
	Email string
}
