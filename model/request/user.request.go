package request

type UserCreateReq struct {
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
	Address string `json:"address" validate:"required"`
	Phone   string `json:"phone" validate:"required"`
}

type UserUpdateReq struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type UserUpdateEmailReq struct {
	Email string `json:"email"`
}
