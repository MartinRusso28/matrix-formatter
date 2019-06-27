package matrixfmt

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/MartinRusso28/matrixfmt/internal/network"
)

//Echo return the matrix as a string in matrix format.
func Echo(file io.Reader) network.Response {
	var body string

	response := network.Response{}

	records, err := readFile(file) 

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

//Invert return the matrix as a string in matrix format where the columns and rows are inverted
func Invert(file io.Reader) network.Response {
	var (
		body     string
		response network.Response
	)

	records, err := readFile(file)

	if err != nil {
		response.Error = errors.New("Cannot read the file")
		response.StatusCode = 500
		return response
	}

	invertedMatrix := invertMatrix(records)

	for _, invertedRow := range invertedMatrix {
		body = fmt.Sprintf("%s%s\n", body, strings.Join(invertedRow, ","))
	}

	response.Body = body

	return response
}

//Flatten return the matrix as a 1 line string, with values separated by commas.
func Flatten(file io.Reader) network.Response {
	var rows []string

	response := network.Response{}

	records, err := readFile(file)

	if err != nil {
		response.Error = err
		response.StatusCode = 500
		return response
	}
	for _, row := range records {
		rows = append(rows, fmt.Sprintf("%s", strings.Join(row, ",")))
	}

	response.Body = strings.Join(rows, ",")

	return response
}

//Sum return the sum of the integers in the matrix
func Sum(file io.Reader) network.Response {
	var sum int

	response := network.Response{}

	records, err := readFile(file) 

	if err != nil {
		response.Error = err
		response.StatusCode = 500
		return response
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

	return response
}

//Multiply return the multiply of the integers in the matrix
func Multiply(file io.Reader) network.Response {
	multiply := 1

	response := network.Response{}

	records, err := readFile(file)

	if err != nil {
		response.Error = err
		response.StatusCode = 500
		return response
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


func readFile(file io.Reader) ([][]string, error){
	records, err := csv.NewReader(file).ReadAll()

	if err != nil {
		return nil, errors.New("Cannot read the file")
	}

	rowsLen := len(records)

	for _, row := range records {
		if len(row) != rowsLen {
			return nil, errors.New("Invalid matrix size, must be square sized")
		}

		for _, cell := range row {
			if cell == ""{
				return nil, errors.New("Invalid cell value, all the cells must have a numeric value")
			}				
		}
	}

	return records, nil
}