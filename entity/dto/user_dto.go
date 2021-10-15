package dto

type LoginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUserDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
