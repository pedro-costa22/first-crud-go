package response

type LoginUserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Age   int8   `json:"age"`
}

type LoginResponse struct {
	Token string            `json:"token"`
	User  LoginUserResponse `json:"user"`
}