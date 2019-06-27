package server

import (
	"github.com/MartinRusso28/matrixfmt/internal/network"
	"github.com/MartinRusso28/matrixfmt/pkg/matrixfmt"
	"github.com/gin-gonic/gin"

	"errors"
	"fmt"
	"io"
)

var (
	response network.Response
)

//GetMainEngine return matrix formatter server.
func GetMainEngine() *gin.Engine {

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	env := environment{}

	router.POST("/echo", env.readFile, env.echo)
	router.POST("/invert", env.readFile, env.invert)
	router.POST("/flatten", env.readFile, env.flatten)
	router.POST("/sum", env.readFile, env.sum)
	router.POST("/multiply", env.readFile, env.multiply)

	return router
}

type environment struct {
}

func (env environment) readFile(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")

	if err != nil {
		response = network.InternalServerError(errors.New("Invalid file"))
		env.respondAndAbort(c, response)
	} else {
		c.Set("file", file)
	}
}

func (env environment) echo(c *gin.Context) {
	file, exists := c.Get("file")

	if !exists {
		response = network.InternalServerError(errors.New("Invalid file"))
		env.respondAndAbort(c, response)
	}

	response = matrixfmt.Echo(file.(io.Reader))

	env.respond(c, response)
}

func (env environment) invert(c *gin.Context) {
	var response network.Response

	file, _, err := c.Request.FormFile("file")

	if err != nil {
		response = network.InternalServerError(errors.New("Invalid file"))
		env.respondAndAbort(c, response)
	}

	response = matrixfmt.Invert(file)

	env.respond(c, response)
}

func (env environment) flatten(c *gin.Context) {
	var response network.Response

	file, _, err := c.Request.FormFile("file")

	if err != nil {
		response = network.InternalServerError(errors.New("Invalid file"))
		env.respondAndAbort(c, response)
	}

	response = matrixfmt.Flatten(file)

	env.respond(c, response)
}

func (env environment) sum(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")

	if err != nil {
		response := network.InternalServerError(errors.New("Invalid file"))
		env.respond(c, response)
	}

	response = matrixfmt.Sum(file)

	env.respond(c, response)
}

func (env environment) multiply(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")

	if err != nil {
		response := network.InternalServerError(errors.New("Invalid file"))
		env.respond(c, response)
	}

	response = matrixfmt.Multiply(file)

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
