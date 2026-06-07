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
		// required: true
		// default: John
		FirstName  string `json:"first_name"`
		// required: true
		// default: Smith
		MiddleName string `json:"middle_name,omitempty"`

		// required: true
		// default: Adams
		LastName   string `json:"last_name"`

		// required: true
		// default: test@example.com
		Email      string `json:"email"`

		// required: true
		// default: 123456
		Password   string `json:"password"`
	}
}