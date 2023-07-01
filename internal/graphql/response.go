package graphql

type UserQueryResponse struct {
	Data DataUser `json:"data"`
	Errors []ErrorMessage `json:"errors"`
}

type UserMutationResponse struct {
	Data DataMutationUser `json:"data"`
	Errors []ErrorMessage `json:"errors"`
}

type DataUser struct {
	Data []User `json:"user"`
}

type DataMutationUser struct {
	Data User `json:"insert_user_one"`
}

type User struct {
	UserID       string `json:"user_id"`
	UserName     string `json:"user_name"`
	UserUsername string `json:"user_username"`
	UserPassword string `json:"user_password"`
}

type ErrorMessage struct {
	Message    string     `json:"message"`
	Extensions Extensions `json:"extensions"`
}

type Extensions struct {
	Code string `json:"code"`
	Path string `json:"path"`
}
