package model

type (
	USER struct {
		ID       int64  `json:"id"`
		UserName string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	CreateUserRequest struct {
		UserName string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	CreateUserResponse struct {
		User USER `json:"user"`
	}

	ReadUserRequest struct {
		UserName string `json:"username"`
	}

	ReadUserResponse struct {
		User USER `json:"user"`
	}

	UpdateUserRequest struct {
		UserName string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	UpdateUserResponse struct {
		User USER `json:"user"`
	}

	DeleteUserRequest struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}

	DeleteUserResponse struct{}
)
