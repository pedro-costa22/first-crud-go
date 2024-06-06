package response

// UserCreateResponse represents the input data for creating a new user.
// @Summary User Input Data
// @Description Structure containing the required fields for return a user.
type UserCreateResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Age   int8   `json:"age"`
}
