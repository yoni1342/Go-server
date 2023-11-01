package types

type SignupRequest struct {
	Input struct {
		Arg struct {
			Email     string `json:"email" graphql:"email"`
			Password  string `json:"password" graphql:"password"`
			FirstName string `json:"first_name" graphql:"first_name"`
			LastName  string `json:"last_name" graphql:"last_name"`
		} `json:"userData" graphql:"userData"`
	} `json:"input" graphql:"input"`
}

type SigninRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type VerifyEmailRequest struct {
	Input struct {
		Arg struct {
			Token string `json:"token" graphql:"token"`
		} `json:"arg1" graphql:"arg1"`
	} `json:"input" graphql:"input"`
}

// type SignupRequest struct {
// 	Input struct {
// 		User struct {
// 			FirstName string `json:"first_name"`
// 			LastName  string `json:"last_name"`
// 			Email     string `json:"email"`
// 			Password  string `json:"password"`
// 		} `json:"user"`
// 	} `json:"input"`
// }

// type SigninRequest struct {
// 	Input struct {
// 		Credential struct {
// 			Email    string `json:"email"`
// 			Password string `json:"password"`
// 		} `json:"credential"`
// 	} `json:"input"`
// }
