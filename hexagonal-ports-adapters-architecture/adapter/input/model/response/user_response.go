package response

type UserResponse struct {
	ID        int32  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int32 `json:"age"`
}
