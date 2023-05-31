package view

type UserRequest struct {
	Username    string   `json:"username"`
	Password    string   `json:"password"`
	Email       string   `json:"email"`
	Phone       string   `json:"phone"`
	Birthdate   string   `json:"birthdate"`
	Perferences []string `json:"perferences"`
}

type UserResponse struct {
	ID           int      `json:"id"`
	Username     string   `json:"username"`
	PasswordHash string   `json:"password"`
	Email        string   `json:"email"`
	Phone        string   `json:"phone"`
	Birthdate    string   `json:"birthdate"`
	Perferences  []string `json:"perferences"`
}

type UserUpdateRequest struct {
	Username   string   `json:"username"`
	Password   string   `json:"password"`
	Prefrences []string `json:"prefrences"`
	ID         int      `json:"id"`
}
