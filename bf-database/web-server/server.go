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
	go queryDatabase(input)
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
	go addNewData(databaseInput)
	context.IndentedJSON(http.StatusOK,)
}
func postRemoveData(context *gin.Context){
	var input databaseInput

	if err :=  context.BindJSON(&input); err != nil{
		return
	}
	go removeData(databaseInput)
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
	setNewCachData(databaseInput)
	persistNewData(databaseInput)
}

func removeData(databaseInput) {
	removeIfCached(databaseInput)
	removePersistedData(databaseInput)
}


func main(){
    router := gin.Default()
	router.POST("/queryData", postQueryDatabaseData)
	router.POST("/createTable", postCreateTable)
	router.POST("/addData", postAddData)
	router.POST("/removeData", postRemoveData)
	router.Run("localhost:8000")
}


// understand go concurrency and channels
// file structure import and exports
// cache mechanism
// btree creation in persistent and cache stores

// how to store btree and data and access to data