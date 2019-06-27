package network

//Response represent a request response.
type Response struct {
	StatusCode int
	Body       string
	Error      error
}
//InternalServerError return a 500 status code with error response.
func InternalServerError(err error) Response {
	response := Response{
		StatusCode: 500,
		Error: err,
	}

	return response
}

//OkResponse return a 200 status code with body response.
func OkResponse(body string) Response {
	response := Response{
		StatusCode: 200,
		Body: body,
	}

	return response
}