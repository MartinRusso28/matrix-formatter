package matrixfmt

import (
	"testing"
	"errors"
	"strconv"
	"github.com/stretchr/testify/assert"
)

func Test_echo_fail_if_records_have_is_non_squared_size(t *testing.T){
	unsquaredMatrix := make([][]string, 2)
	unsquaredMatrix[0] = make([]string,1)

	response := Echo(unsquaredMatrix)

	assert.Equal(t, 500, response.StatusCode)
	assert.Equal(t, errors.New("Invalid matrix size, must be square sized"), response.Error)
}

func Test_echo_fail_if_records_have_a_non_numeric_cell(t *testing.T){
	nonNumMatrix := make([][]string, 2)
	nonNumMatrix[0] = make([]string, 2)
	nonNumMatrix[1] = make([]string, 2)

	for i, row := range nonNumMatrix {
		for j, _ := range row {
			nonNumMatrix[i][j] = "1"
		}
	}

	nonNumMatrix[0][0] = "a"

	response := Echo(nonNumMatrix)

	assert.Equal(t, 500, response.StatusCode)
	assert.Equal(t, errors.New("Invalid cell value, all the cells must have a numeric value"), response.Error)
}

func Test_echo_success_if_have_valid_record(t *testing.T){
	nonNumMatrix := make([][]string, 2)
	nonNumMatrix[0] = make([]string, 2)
	nonNumMatrix[1] = make([]string, 2)

	for i, row := range nonNumMatrix {
		for j, _ := range row {
			num := strconv.Itoa(i)
			nonNumMatrix[i][j] = num
		}
	}

	response := Echo(nonNumMatrix)

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "0,0\n1,1\n", response.Body)
}

func Test_Invert_fail_if_records_have_is_non_squared_size(t *testing.T){
	unsquaredMatrix := make([][]string, 2)
	unsquaredMatrix[0] = make([]string,1)

	response := Invert(unsquaredMatrix)

	assert.Equal(t, 500, response.StatusCode)
	assert.Equal(t, errors.New("Invalid matrix size, must be square sized"), response.Error)
}

func Test_Invert_fail_if_records_have_a_non_numeric_cell(t *testing.T){
	nonNumMatrix := make([][]string, 2)
	nonNumMatrix[0] = make([]string, 2)
	nonNumMatrix[1] = make([]string, 2)

	for i, row := range nonNumMatrix {
		for j, _ := range row {
			nonNumMatrix[i][j] = "1"
		}
	}

	nonNumMatrix[0][0] = "a"

	response := Invert(nonNumMatrix)

	assert.Equal(t, 500, response.StatusCode)
	assert.Equal(t, errors.New("Invalid cell value, all the cells must have a numeric value"), response.Error)
}

func Test_Invert_success_if_have_valid_record(t *testing.T){
	nonNumMatrix := make([][]string, 2)
	nonNumMatrix[0] = make([]string, 2)
	nonNumMatrix[1] = make([]string, 2)

	for i, row := range nonNumMatrix {
		for j, _ := range row {
			num := strconv.Itoa(i)
			nonNumMatrix[i][j] = num
		}
	}

	response := Invert(nonNumMatrix)

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "0,1\n0,1\n", response.Body)
}

func Test_Flatten_fail_if_records_have_is_non_squared_size(t *testing.T){
	unsquaredMatrix := make([][]string, 2)
	unsquaredMatrix[0] = make([]string,1)

	response := Flatten(unsquaredMatrix)

	assert.Equal(t, 500, response.StatusCode)
	assert.Equal(t, errors.New("Invalid matrix size, must be square sized"), response.Error)
}

func Test_Flatten_fail_if_records_have_a_non_numeric_cell(t *testing.T){
	nonNumMatrix := make([][]string, 2)
	nonNumMatrix[0] = make([]string, 2)
	nonNumMatrix[1] = make([]string, 2)

	for i, row := range nonNumMatrix {
		for j, _ := range row {
			nonNumMatrix[i][j] = "1"
		}
	}

	nonNumMatrix[0][0] = "a"

	response := Flatten(nonNumMatrix)

	assert.Equal(t, 500, response.StatusCode)
	assert.Equal(t, errors.New("Invalid cell value, all the cells must have a numeric value"), response.Error)
}

func Test_Flatten_success_if_have_valid_record(t *testing.T){
	nonNumMatrix := make([][]string, 2)
	nonNumMatrix[0] = make([]string, 2)
	nonNumMatrix[1] = make([]string, 2)

	for i, row := range nonNumMatrix {
		for j, _ := range row {
			num := strconv.Itoa(i)
			nonNumMatrix[i][j] = num
		}
	}

	response := Flatten(nonNumMatrix)

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "0,0,1,1", response.Body)
}

func Test_Sum_fail_if_records_have_is_non_squared_size(t *testing.T){
	unsquaredMatrix := make([][]string, 2)
	unsquaredMatrix[0] = make([]string,1)

	response := Sum(unsquaredMatrix)

	assert.Equal(t, 500, response.StatusCode)
	assert.Equal(t, errors.New("Invalid matrix size, must be square sized"), response.Error)
}

func Test_Sum_fail_if_records_have_a_non_numeric_cell(t *testing.T){
	nonNumMatrix := make([][]string, 2)
	nonNumMatrix[0] = make([]string, 2)
	nonNumMatrix[1] = make([]string, 2)

	for i, row := range nonNumMatrix {
		for j, _ := range row {
			nonNumMatrix[i][j] = "1"
		}
	}

	nonNumMatrix[0][0] = "a"

	response := Sum(nonNumMatrix)

	assert.Equal(t, 500, response.StatusCode)
	assert.Equal(t, errors.New("Invalid cell value, all the cells must have a numeric value"), response.Error)
}

func Test_Sum_success_if_have_valid_record(t *testing.T){
	nonNumMatrix := make([][]string, 2)
	nonNumMatrix[0] = make([]string, 2)
	nonNumMatrix[1] = make([]string, 2)

	for i, row := range nonNumMatrix {
		for j, _ := range row {
			num := strconv.Itoa(i)
			nonNumMatrix[i][j] = num
		}
	}

	response := Sum(nonNumMatrix)

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "2", response.Body)
}

func Test_Multiply_fail_if_records_have_is_non_squared_size(t *testing.T){
	unsquaredMatrix := make([][]string, 2)
	unsquaredMatrix[0] = make([]string,1)

	response := Multiply(unsquaredMatrix)

	assert.Equal(t, 500, response.StatusCode)
	assert.Equal(t, errors.New("Invalid matrix size, must be square sized"), response.Error)
}

func Test_Multiply_fail_if_records_have_a_non_numeric_cell(t *testing.T){
	nonNumMatrix := make([][]string, 2)
	nonNumMatrix[0] = make([]string, 2)
	nonNumMatrix[1] = make([]string, 2)

	for i, row := range nonNumMatrix {
		for j, _ := range row {
			nonNumMatrix[i][j] = "1"
		}
	}

	nonNumMatrix[0][0] = "a"

	response := Multiply(nonNumMatrix)

	assert.Equal(t, 500, response.StatusCode)
	assert.Equal(t, errors.New("Invalid cell value, all the cells must have a numeric value"), response.Error)
}

func Test_Multiply_success_if_have_valid_record(t *testing.T){
	nonNumMatrix := make([][]string, 2)
	nonNumMatrix[0] = make([]string, 2)
	nonNumMatrix[1] = make([]string, 2)

	for i, row := range nonNumMatrix {
		for j, _ := range row {
			nonNumMatrix[i][j] = "2"
		}
	}

	response := Multiply(nonNumMatrix)

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "16", response.Body)
}