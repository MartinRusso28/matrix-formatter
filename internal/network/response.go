package network

//Response represent a request response.
type Response struct {
	StatusCode int `valid:"required"`
	Body       string
	Error      error
}
//InternalServerError response.
func InternalServerError(err error) Response {
	response := Response{
		StatusCode: 500,
		Error: err,
	}

	return response
}