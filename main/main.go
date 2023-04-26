package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type process struct {
	ID    string `json:"id"`
	Model string `json:"model"`
	Meta  string `json:"meta"`
	Step  string `json:"step"`
}

var processes = []process{
	{ID: "1", Model: "", Meta: "", Step: ""},
	{ID: "2", Model: "", Meta: "", Step: ""},
	{ID: "3", Model: "", Meta: "", Step: ""},
}

func main() {
	router := gin.Default()
	router.GET("/process", getProcesses)

	router.Run("localhost:8080")
}

func getProcesses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, processes)
}
