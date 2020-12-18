package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/add", sum)
	router.POST("/multiply", multiply)
	router.POST("/substruct", substruction)
	router.POST("/division", division)

	router.Run(":3000")
}

// Request struct
type Request struct {
	Num1 int64 `json:"num1"`
	Num2 int64 `json:"num2"`
}

// Response struct
type Response struct {
	Result int64  `json:"result"`
	Error  string `json:"error"`
}

func sum(c *gin.Context) {
	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Result: 0,
			Error:  "Error in data binding",
		})
		return
	}

	result := req.Num1 + req.Num2
	c.JSON(http.StatusOK, Response{
		Result: result,
		Error:  "",
	})
	return
}

func multiply(c *gin.Context) {
	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Result: 0,
			Error:  "Error in data binding",
		})
		return
	}

	result := req.Num1 * req.Num2
	c.JSON(http.StatusOK, Response{
		Result: result,
		Error:  "",
	})
	return
}

func substruction(c *gin.Context) {
	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Result: 0,
			Error:  "Error in data binding",
		})
		return
	}

	result := req.Num1 - req.Num2
	c.JSON(http.StatusOK, Response{
		Result: result,
		Error:  "",
	})
	return
}
func division(c *gin.Context) {
	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Result: 0,
			Error:  "Error in data binding",
		})
		return
	}

	result := req.Num1 / req.Num2
	c.JSON(http.StatusOK, Response{
		Result: result,
		Error:  "",
	})
	return
}
