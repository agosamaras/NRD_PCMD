package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func storeData(ctx *gin.Context) {
	// Validate input
	var input []interface{} //TODO(user): you should bind to input the type that suits you best.
	if err := ctx.ShouldBindJSON(&input); err != nil {
		log.Error("Unable to handle request --> ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Info("Storing/Modifying received Data")
	// TODO(user): here you can modify to csv your json data and store
	// them accordingly
	ctx.JSON(http.StatusCreated, input) // or c.IndentedJSON(http.StatusNotFound, gin.H{"message": "data stored."})
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "This PCMD server receives data in JSON format.",
		})
	})
	r.POST("/", storeData)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080") the same as r.Run("localhost:8080")
}
