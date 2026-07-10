package request

type UserRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int32  `json:"age"`
}
