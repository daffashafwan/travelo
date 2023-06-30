package graphql

type UserResponse struct {
	User []User `json:"user"`
}

type User struct {
	UserID       string `json:"user_id"`
	UserName     string `json:"user_name"`
	UserUsername string `json:"user_username"`
	UserPassword string `json:"user_password"`
}
