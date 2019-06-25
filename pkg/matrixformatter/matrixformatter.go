package matrixformatter

import (
	"encoding/csv"
	"fmt"
	"strings"
	"io"
	"github.com/MartinRusso28/matrix-formatter/internal/network"

)

//Echo return the matrix as a string in matrix format.
func Echo(file io.Reader) network.Response {
	var body string

	response := network.Response{}

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		response.Error = err
		response.StatusCode = 500
		// w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return response
	}
	for _, row := range records {
		body = fmt.Sprintf("%s%s\n", body, strings.Join(row, ","))
	}

	response.Body = body

	return response
}

//Invert return the matrix as a string in matrix format where the columns and rows are inverted
func Invert(file io.Reader) network.Response {
	var body string

	response := network.Response{}

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		response.Error = err
		response.StatusCode = 500
		return response
	}
	for _, row := range records {
		body = fmt.Sprintf("%s%s\n", body, strings.Join(row, ","))
	}

	response.Body = body

	return response
}

//Flatten return the matrix as a 1 line string, with values separated by commas.
func Flatten(file io.Reader) network.Response {
	var body string

	response := network.Response{}

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		response.Error = err
		response.StatusCode = 500
		return response
	}
	for _, row := range records {
		body = fmt.Sprintf("%s %s", body, strings.Join(row, ","))
	}

	response.Body = body

	return response
}