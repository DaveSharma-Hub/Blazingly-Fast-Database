package server

import (
    "net/http"
    "github.com/gin-gonic/gin"
	// "github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/cache/lruCache"
	// "github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/cache"
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/dataCacheClient"
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/types"
	// "github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/persistentStore"
	"fmt"
	"encoding/json"
)

type FunctionWrapperType func(*gin.Context, dataCacheClient.DataCacheExecutionType)

type PostQueryInputType struct{
	TableName string `json:"table_name" binding:"required"`
	PartitionKey string `json:"partition_key" binding:"required"`
}

// type databaseOutput struct{
// 	Output string `json:"output"`
// 	Data string `json:"data"`
// }

// var tmpData = []databaseOutput{
// 	{Output: "Blue Train", Data: "John Coltrane"},
// 	{Output: "Blue Train", Data: "John Coltrane"},
// }

func postQueryDatabaseData(c *gin.Context, executeFn dataCacheClient.DataCacheExecutionType){
	var inputData PostQueryInputType

	if err:= c.ShouldBindJSON(&inputData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	var returnData globalTypes.Payload = executeFn(inputData.TableName, inputData.PartitionKey, globalTypes.CreateEmptyPayload())
	jsonResult, err := json.Marshal(returnData.Item)
	if err!=nil {
		fmt.Println("ERROR")
	}
	fmt.Println(returnData)
	c.JSON(http.StatusOK, string(jsonResult))
}

func CreateFunctionWrapper(inputFn FunctionWrapperType, client dataCacheClient.DataCacheClientReturnType, functionName string)gin.HandlerFunc{
    return func (c *gin.Context) {
        inputFn(c, client[functionName])
    }
}

func InitServer(client dataCacheClient.DataCacheClientReturnType)*gin.Engine{
    router := gin.Default()
	//router.GET("/test",CreateFunctionWrapper(testGet,client, "QueryData"))
	// router.POST("/test",testPost)
	router.POST("/queryData", CreateFunctionWrapper(postQueryDatabaseData, client, "QueryData"))
	// router.POST("/createTable", postCreateTable)
	// router.POST("/addData", postAddData)
	// router.POST("/removeData", postRemoveData)
	
	return router
}

func RunServer(router *gin.Engine){
	router.Run("localhost:8000")
}
