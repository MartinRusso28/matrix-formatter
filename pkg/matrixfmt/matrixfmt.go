package matrixfmt

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/MartinRusso28/matrixfmt/internal/network"
)

//Echo return the matrix as a string in matrix format.
func Echo(records [][]string) network.Response {
	var body string

	response := network.Response{}

	err := validateRecords(records) 

	if err != nil {
		return network.InternalServerError(err)
	}

	for _, row := range records {
		body = fmt.Sprintf("%s%s\n", body, strings.Join(row, ","))
	}

	response.Body = body
	response.StatusCode = 200

	return response
}

//Invert return the matrix as a string in matrix format where the columns and rows are inverted
func Invert(records [][]string) network.Response {
	var (
		body     string
		response network.Response
	)

	err := validateRecords(records)

	if err != nil {
		return network.InternalServerError(err)
	}

	invertedMatrix := invertMatrix(records)

	for _, invertedRow := range invertedMatrix {
		body = fmt.Sprintf("%s%s\n", body, strings.Join(invertedRow, ","))
	}

	response.Body = body
	response.StatusCode = 200

	return response
}

//Flatten return the matrix as a 1 line string, with values separated by commas.
func Flatten(records [][]string) network.Response {
	var rows []string

	response := network.Response{}

	err := validateRecords(records)

	if err != nil {
		return network.InternalServerError(err)
	}

	for _, row := range records {
		rows = append(rows, fmt.Sprintf("%s", strings.Join(row, ",")))
	}

	response.Body = strings.Join(rows, ",")
	response.StatusCode = 200

	return response
}

//Sum return the sum of the integers in the matrix
func Sum(records [][]string) network.Response {
	var sum int

	response := network.Response{}

	err := validateRecords(records) 

	if err != nil {
		return network.InternalServerError(err)
	}

	for _, row := range records {
		for _, cell := range row {
			num, err := strconv.Atoi(cell)

			if err != nil {
				response := network.InternalServerError(fmt.Errorf("Invalid cell: %s", cell))
				return response
			}

			sum += num
		}
	}

	response.Body = strconv.Itoa(sum)
	response.StatusCode = 200

	return response
}

//Multiply return the multiply of the integers in the matrix
func Multiply(records [][]string) network.Response {
	multiply := 1

	response := network.Response{}

	err := validateRecords(records)

	if err != nil {
		return network.InternalServerError(err)
	}

	for _, row := range records {
		for _, cell := range row {
			num, err := strconv.Atoi(cell)

			if err != nil {
				response := network.InternalServerError(fmt.Errorf("Invalid cell: %s", cell))
				return response
			}

			multiply *= num
		}
	}

	response.Body = strconv.Itoa(multiply)
	response.StatusCode = 200

	return response
}

func invertMatrix(matrix [][]string) [][]string {
	invertedMatrix := initializeMatrix(len(matrix))

	for j, row := range matrix {
		for k, cell := range row {
			invertedMatrix[k][j] = cell
		}
	}

	return invertedMatrix
}

func initializeMatrix(len int) [][]string {
	matrix := make([][]string, len)

	for i := 0; i < len; i++ {
		matrix[i] = make([]string, len)
	}

	return matrix
}

func validateRecords(records [][]string) (error){
	rowsLen := len(records)

	for _, row := range records {
		if len(row) != rowsLen {
			return errors.New("Invalid matrix size, must be square sized")
		}

		for _, cell := range row {
			_, err := strconv.Atoi(cell)

			if (err != nil) {
				return errors.New("Invalid cell value, all the cells must have a numeric value")
			}				
		}
	}

	return nil
}