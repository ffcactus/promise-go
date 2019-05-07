package dto

type PostAccountRequest struct {
	UserName string `json:"UserName"`
	Password string `json:"Password"`
	RoleId   string `json:"RoleId"`
}
