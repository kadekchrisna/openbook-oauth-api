package users

type (
	User struct {
		Id        int64  `json:"id"`
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}
	UserLoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)
