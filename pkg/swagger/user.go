package swagger

type UserRequest struct {
	Name    string `json:"name" validate:"required" example:"user"`
	Password string `json:"password" validate:"required" example:"password"`
}
