package request

type CreateUserAccountRequest struct {
	ID   string `json:"id"`
	Pass string `json:"password"`
}
