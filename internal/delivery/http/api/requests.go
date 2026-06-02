package api_requests

// swagger:parameters createUser
type createUserParams struct {
    // in: body
    // required: true
    Body struct {
        FirstName string `json:"first_name"`
		MiddleName string `json:"middle_name,omitempty"`
        LastName  string `json:"last_name"`
        Email     string `json:"email"`
        Password  string `json:"password"`
    }
}

// swagger:parameters registerUser
type registerUserParams struct {
	// in: body
	// required: true
	Body struct {
		FirstName  string `json:"first_name"`
		MiddleName string `json:"middle_name,omitempty"`
		LastName   string `json:"last_name"`
		Email      string `json:"email"`
		Password   string `json:"password"`
	}
}