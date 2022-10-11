package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello/:name", helloHandler)
	r.Run()
}

type HelloMessage struct {
	name string
}

func (m HelloMessage) String() string {
	return fmt.Sprintf("Hello, %s", m.name)
}

func helloHandler(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"message": HelloMessage{name: name}.String(),
	})
}
