# Matrix formatter

For fun the server:

    1- cd cmd/
    2- go run main.go

## Server: localhost:8080

### Methods (output using the csv example)

1. /echo
    - Return the matrix as a string in matrix format.
    
    ```
    output:
    1,2,3
    4,5,6
    7,8,9
    ``` 
2. /invert
    - Return the matrix as a string in matrix format where the columns and rows are inverted
    ```
    output:
    1,4,7
    2,5,8
    3,6,9
    ``` 
3. /flatten
    - Return the matrix as a 1 line string, with values separated by commas.
    ```
    output:
    1,2,3,4,5,6,7,8,9
    ``` 
4. /sum
    - Return the sum of the integers in the matrix
    ```
    output:
    45
    ``` 
5. /multiply
    - Return the product of the integers in the matrix
    ```
    output:
    362880
    ``` 

## Send the .csv file in the body with the 'file' param.

The input file to these functions is a matrix, of any dimension where the number of rows are equal to the number of columns (square). Each value is an integer, and there is no header row. matrix.csv is example valid input.

