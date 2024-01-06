package server

import (
    "net/http"
    "github.com/gin-gonic/gin"
	// "github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/cache/lruCache"
	// "github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/cache"
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/dataCacheClient"
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/types"
	// "github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/persistentStore"
)


type FunctionWrapperType func(*gin.Context, dataCacheClient.DataCacheExecutionType)

type databaseOutput struct{
	Output string `json:"output"`
	Data string `json:"data"`
}

var tmpData = []databaseOutput{
	{Output: "Blue Train", Data: "John Coltrane"},
	{Output: "Blue Train", Data: "John Coltrane"},
}

func testGet(c *gin.Context, executeFn dataCacheClient.DataCacheExecutionType){
	returnData := executeFn("1", globalTypes.CreateEmptyPayload())
	fmt.Println(returnData)	
	c.IndentedJSON(http.StatusOK, returnData)
}

func CreateFunctionWrapper(inputFn FunctionWrapperType, client dataCacheClient.DataCacheClientReturnType, functionName string)gin.HandlerFunc{
    return func (c *gin.Context) {
        inputFn(c, client[functionName])
    }
}

func InitServer(client dataCacheClient.DataCacheClientReturnType)*gin.Engine{
    router := gin.Default()
	router.GET("/test",CreateFunctionWrapper(testGet,client, "GetData"))
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
