package network

import (
	"github.com/asaskevich/govalidator"
)

//Response represent a request response.
type Response struct {
	StatusCode int `valid:"required"`
	Body       string
	Error      error
}

//Valid response struct.
func (r Response) Valid() bool {
	valid, _ := govalidator.ValidateStruct(r)
	return valid
}

//InternalServerError response.
func InternalServerError(err error) Response {
	response := Response{
		StatusCode: 500,
		Error: err,
	}

	return response
}