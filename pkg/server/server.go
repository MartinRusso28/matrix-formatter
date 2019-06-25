package server

import (
	"github.com/gin-gonic/gin"
	"github.com/MartinRusso28/matrix-formatter/pkg/matrixformatter"
	"github.com/MartinRusso28/matrix-formatter/internal/network"

	"fmt"
	"errors"
	"io"
)

var (
	response network.Response
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

//GetMainEngine return money converter's server.
func GetMainEngine() *gin.Engine {

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	env := environment{}

	router.POST("/echo",env.readFile, env.echo)
	router.POST("/invert", env.readFile, env.invert)
	router.POST("/flatten", env.readFile, env.flatten)

	return router
}

type environment struct {
}

func (env environment) readFile(c *gin.Context){
	file,_, err := c.Request.FormFile("file")

	if err != nil {
		response.StatusCode = 500
		response.Error = errors.New("Invalid file")
		env.respondAndAbort(c, response)
	} else {
		c.Set("file", file)
	}
}

func (env environment) echo(c *gin.Context) {
	file, exists := c.Get("file")

	if !exists {
		response.StatusCode = 500
		response.Error = errors.New("Invalid file")
	}

	response = matrixformatter.Echo(file.(io.Reader))
	
	env.respond(c, response)
}

func (env environment) invert(c *gin.Context) {
	var response network.Response

	file,_, err := c.Request.FormFile("file")

	if err != nil {
		response.StatusCode = 500
		response.Error = errors.New("Invalid file")
	} else {
		response = matrixformatter.Invert(file)
	}
	
	env.respond(c, response)
}

func (env environment) flatten(c *gin.Context) {
	var response network.Response

	file,_, err := c.Request.FormFile("file")

	if err != nil {
		response.StatusCode = 500
		response.Error = errors.New("Invalid file")
	} else {
		response = matrixformatter.Flatten(file)
	}
	
	env.respond(c, response)
}


func (env environment) respond(c *gin.Context, response network.Response) {
	obj := gin.H{}

	if response.Error != nil {
		obj["error"] = response.Error.Error()
		c.JSON(response.StatusCode, obj)
	}
	if response.Body != "" {
		fmt.Fprint(c.Writer, response.Body)
	}
}

func (env environment) respondAndAbort(c *gin.Context, response network.Response) {
	env.respond(c, response)
	c.Abort()
}

// func main() {
// 	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
// 		file, _, err := r.FormFile("file")
// 		if err != nil {
// 			w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
// 			return
// 		}
// 		defer file.Close()
// 		records, err := csv.NewReader(file).ReadAll()
// 		if err != nil {
// 			w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
// 			return
// 		}
// 		var response string
// 		for _, row := range records {
// 			response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
// 		}
// 		fmt.Fprint(w, response)
// 	})
// 	http.ListenAndServe(":8080", nil)
// }
