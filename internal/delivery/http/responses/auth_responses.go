package http_responses

import "github.com/Ozenkol/rbk-go-final/internal/application/service"

// swagger:response loginUserResponse
type loginUserResponse struct {
	// in: body
	Body service.TokenPair
}

// swagger:response registerUserResponse
type registerUserResponse struct {
	// in: body
	Body struct {
		ID    string `json:"id"`
		Email string `json:"email"`
	}
}
