package repository

type UserRepository interface {
	CreateUsersAccount(id, pass, token string) error
	RegisterUsersInfo(id, pass, token string) error
}
