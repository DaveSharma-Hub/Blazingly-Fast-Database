package server

import (
    "net/http"
    "github.com/gin-gonic/gin"
	// "github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/cache/lruCache"
	// "github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/cache"
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/dataCacheClient"
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/types"
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/server/utils"
	// "github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/persistentStore"
	"fmt"
	"encoding/json"
)



// type databaseOutput struct{
// 	Output string `json:"output"`
// 	Data string `json:"data"`
// }

// var tmpData = []databaseOutput{
// 	{Output: "Blue Train", Data: "John Coltrane"},
// 	{Output: "Blue Train", Data: "John Coltrane"},
// }

func postQueryDatabaseData(c *gin.Context, executeFn dataCacheClient.DataCacheExecutionType){
	var inputData utils.PostQueryInputType
	
	if err:= c.ShouldBindJSON(&inputData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	var returnData globalTypes.Payload = executeFn(inputData.TableName, inputData.PartitionKey, globalTypes.CreateEmptyPayload())
	jsonResult, err := json.Marshal(returnData.Item)
	if err!=nil {
		fmt.Println("ERROR")
	}
	c.JSON(http.StatusOK, string(jsonResult))
}

func postSetDatabaseData(c *gin.Context, executeFn dataCacheClient.DataCacheExecutionType){
	var inputData utils.PostSetDataInputType
	// inputData.DataPayload = make(map[string] globalTypes.AtomicItem)
	
	if err:= c.ShouldBindJSON(&inputData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newPayload := utils.GetPayloadFromPostSetDataInput(inputData)
	executeFn(inputData.TableName, inputData.PartitionKey, newPayload)

	jsonResult, err := json.Marshal(inputData)

	if err!=nil {
		fmt.Println("ERROR")
	}
	c.JSON(http.StatusOK, string(jsonResult))
}
func postUpdateDatabaseData(c *gin.Context, executeFn dataCacheClient.DataCacheExecutionType){
	var inputData utils.PostSetDataInputType
	// inputData.DataPayload = make(map[string] globalTypes.AtomicItem)
	
	if err:= c.ShouldBindJSON(&inputData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newPayload := utils.GetPayloadFromPostSetDataInput(inputData)
	executeFn(inputData.TableName, inputData.PartitionKey, newPayload)

	jsonResult, err := json.Marshal(inputData)

	if err!=nil {
		fmt.Println("ERROR")
	}
	c.JSON(http.StatusOK, string(jsonResult))
}

func CreateFunctionWrapper(inputFn utils.FunctionWrapperType, client dataCacheClient.DataCacheClientReturnType, functionName string)gin.HandlerFunc{
    return func (c *gin.Context) {
		// contextCopy := c.Copy()
        inputFn(c, client[functionName])
    }
}

func InitServer(client dataCacheClient.DataCacheClientReturnType)*gin.Engine{
    router := gin.Default()
	//router.GET("/test",CreateFunctionWrapper(testGet,client, "QueryData"))
	// router.POST("/test",testPost)
	router.POST("/queryData", CreateFunctionWrapper(postQueryDatabaseData, client, "QueryData"))
	// router.POST("/createTable", postCreateTable)
	router.POST("/addData", CreateFunctionWrapper(postSetDatabaseData, client, "SetData"))
	router.POST("/updateData", CreateFunctionWrapper(postUpdateDatabaseData, client, "UpdateData"))
	// router.POST("/removeData", postRemoveData)
	
	return router
}

func RunServer(router *gin.Engine){
	router.Run("localhost:8000")
}
