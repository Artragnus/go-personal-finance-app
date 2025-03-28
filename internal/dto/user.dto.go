package dto

type UserCreateRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserCreateReponse struct {
	Name string `json:"name"`
}
