package server

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type databaseOutput struct{
	Output string `json:"output"`
	Data string `json:"data"`
}

var tmpData = []databaseOutput{
	{Output: "Blue Train", Data: "John Coltrane"},
	{Output: "Blue Train", Data: "John Coltrane"},
}

func testGet(c *gin.Context){
	c.IndentedJSON(http.StatusOK, tmpData)
}

func InitServer()*gin.Engine{
    router := gin.Default()
	router.GET("/test",testGet)
	// router.POST("/test",testPost)
	// router.POST("/queryData", postQueryDatabaseData)
	// router.POST("/createTable", postCreateTable)
	// router.POST("/addData", postAddData)
	// router.POST("/removeData", postRemoveData)
	
	return router
}

func RunServer(router *gin.Engine){
	router.Run("localhost:8000")
}
