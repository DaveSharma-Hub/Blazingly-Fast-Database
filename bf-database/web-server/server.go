package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type databaseInput struct{
	InputCommandType string `json:"inputcommandtype"`
	Key string `json:"key"`
	TableName string `json:"tablename"`
	Data string `json:"data"`
}
// inputcommand eg. GET <header/key value> FROM <table name>
// INSERT <key value> INTO <table name> DATA

type databaseOutput struct{
	Output string `json:"output"`
	Data string `json:"data"`
}

type tableHeader struct{
	tableHeaderName string
	tableHeaderType string
}

var tmpData = []databaseOutput{
	{Output: "Blue Train", Data: "John Coltrane"},
	{Output: "Blue Train", Data: "John Coltrane"},
}

func postQueryDatabaseData(context *gin.Context){
	var input databaseInput

	if err :=  context.BindJSON(&input); err != nil{
		return
	}
	go queryDatabaseWorker(input)
	context.IndentedJSON(http.StatusOK,)
}
func postCreateTable(context *gin.Context){
	var input databaseInput

	if err :=  context.BindJSON(&input); err != nil{
		return
	}

	go createTableWorker(input)
	context.IndentedJSON(http.StatusOK,)
}
func postAddData(context *gin.Context){
	var input databaseInput

	if err :=  context.BindJSON(&input); err != nil{
		return
	}
	go addNewDataWorker(databaseInput)
	context.IndentedJSON(http.StatusOK,)
}
func postRemoveData(context *gin.Context){
	var input databaseInput

	if err :=  context.BindJSON(&input); err != nil{
		return
	}
	go removeDataWorker(databaseInput)
	context.IndentedJSON(http.StatusOK,)
}

func createTableWorker(databaseInput input){
	// create table in file system
	persistCreateTable(databaseInput)
}

func queryDatabase(databaseInput input){
	var cacheData = getCachedData(input)
	if(!cacheData){
		return getPersistedData(input)
	}else{
		return cacheData
	}
}

func addNewData(databaseInput) {
	setNewCacheData(databaseInput)
	persistNewData(databaseInput)
}

func removeData(databaseInput) {
	removeIfCached(databaseInput)
	removePersistedData(databaseInput)
}


func testGet(c *gin.Context){
	c.IndentedJSON(http.StatusOK, tmpData)
}


type testType struct{
	input string `json:"input"`
}

func testPost(c *gin.Context){
	var testValue testType

	// Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&testValue); err != nil {
        return
    }
	newValue := databaseOutput{Output:testValue.input, Data:"new data"}

    // Add the new album to the slice.
    tmpData = append(tmpData, newValue)
    c.IndentedJSON(http.StatusCreated, tmpData)
}

func main(){
    router := gin.Default()
	router.GET("/test",testGet)
	router.POST("/test",testPost)
	// router.POST("/queryData", postQueryDatabaseData)
	// router.POST("/createTable", postCreateTable)
	// router.POST("/addData", postAddData)
	// router.POST("/removeData", postRemoveData)
	router.Run("localhost:8000")
}


// understand go concurrency and channels
// file structure import and exports
// cache mechanism
// btree creation in persistent and cache stores

// how to store btree and data and access to data