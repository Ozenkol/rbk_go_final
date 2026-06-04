package http_responses

// swagger:response createUserResponse
type createUserResponse struct {
    // in: body
    Body struct {
        Message string `json:"message"`
    }
}

// swagger:response getUserResponse
type getUserResponse struct {
    // in: body
    Body string `json:"body"`
}

// swagger:response errorResponse
type errorResponse struct {
    // in: body
    Body struct {
        Error string `json:"error"`
    }
}
