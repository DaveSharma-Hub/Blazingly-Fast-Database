package persistentStoreClient

import (
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/types"
	// "github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/cache/binaryTree"
	// "time"
	"fmt"
)

type TableMetaDataType struct{
	TableAttributes globalTypes.TableSchema
}

// here is where you can optimize using binary tree's etc
type TableDataType struct{
	Data map[string] *globalTypes.Payload	// Array of payload that matches schema
}

// type TableDataType struct{
// 	Items map[string] *ItemTableDataType
// }

type MetaAndDataEncapsulation struct{
	TableData *TableDataType // map of Table Data Types with key being table name
	TableMetaData *TableMetaDataType // meta data of Tables
}

type TableEncapsulation struct{
	TableInformation map[string] *MetaAndDataEncapsulation
}

func CreateTableMetaData(schema globalTypes.TableSchema)*TableMetaDataType{
	var metaData TableMetaDataType
	metaData.TableAttributes = schema
	return &metaData
}

func CreateTableData()*TableDataType{
	tableData := TableDataType{Data:make(map[string] *globalTypes.Payload)}
	return &tableData
}

func CreateMetaAndDataEncapsulation(schema globalTypes.TableSchema)*MetaAndDataEncapsulation{
	var encapslation MetaAndDataEncapsulation = MetaAndDataEncapsulation{TableData:CreateTableData(),TableMetaData:CreateTableMetaData(schema)}
	return &encapslation
}

func CreateTableEncapsulation()*TableEncapsulation{
	var allTableData TableEncapsulation
	allTableData.TableInformation = make(map[string] *MetaAndDataEncapsulation)
	return &allTableData
}

func AddTable(schema globalTypes.TableSchema,info *TableEncapsulation){
	metaEncapsulation := CreateMetaAndDataEncapsulation(schema)
	info.TableInformation[schema.TableName] = metaEncapsulation
	fmt.Println(schema.TableName)
}


func tempStoreData()*TableEncapsulation{
	var allTableData *TableEncapsulation = CreateTableEncapsulation()

	CreateTable("Users",[][]string{{"id","string"},{"name","string"},{"age","integer"},{"occupation","string"}},allTableData)
	CreateTable("Locations",[][]string{{"id","string"},{"city","string"},{"country","string"}},allTableData)

	SetData("Users","First",globalTypes.CreatePayload([][]string{{"id","1","string"},{"name","John","string"},{"age","24","integer"},{"occupation","Engineer","string"}}),allTableData)
	SetData("Users","Second",globalTypes.CreatePayload([][]string{{"id","2","string"},{"name","Bob","string"},{"age","74","integer"},{"occupation","Plumber","string"}}),allTableData)
	SetData("Users","Third",globalTypes.CreatePayload([][]string{{"id","3","string"},{"name","Kelly","string"},{"age","44","integer"},{"occupation","Financy","string"}}),allTableData)

	return allTableData
}

func InitPersistentStoreClient()*TableEncapsulation{
	return tempStoreData()
}

func GetData(tableName string, key string, allTableData *TableEncapsulation) globalTypes.Payload{
	// time.Sleep(1000000000)
	// return globalTypes.CreatePayload([][]string{{"id","1","string"},{"name","John","string"},{"age","14","integer"},{"occupation","Engineer","string"}})
	_, ok := allTableData.TableInformation[tableName]
	if ok {
		if allTableData.TableInformation[tableName].TableData != nil {
			if allTableData.TableInformation[tableName].TableData.Data != nil {
				_, okAgain := allTableData.TableInformation[tableName].TableData.Data[key]
				if okAgain {
					return *allTableData.TableInformation[tableName].TableData.Data[key]
				}
			}
		}
	}
	return globalTypes.CreateEmptyPayload();
}

func SetData(tableName string, key string, value globalTypes.Payload, allTableData *TableEncapsulation){
	_, ok := allTableData.TableInformation[tableName]
	if ok {
		if allTableData.TableInformation[tableName].TableData != nil {
			if allTableData.TableInformation[tableName].TableData.Data != nil {
				_, ok := allTableData.TableInformation[tableName].TableData.Data[key]
				if ok {
					// do nothing if it already exists	
				} else {
					allTableData.TableInformation[tableName].TableData.Data[key] = &value
					
				}
			}
		}
	}
}

func UpdateData(tableName string, key string, value globalTypes.Payload, allTableData *TableEncapsulation){
	_, ok := allTableData.TableInformation[tableName]
	if ok {
		if allTableData.TableInformation[tableName].TableData != nil {
			if allTableData.TableInformation[tableName].TableData.Data != nil {
				_,ok := allTableData.TableInformation[tableName].TableData.Data[key]
				if ok {
					// allTableData.TableInformation[tableName].TableData.Data[key] = &value
					// update data 
				}
			}
		}
	}
}


func CreateTable(tableName string, tableSchema[][]string, allTableData *TableEncapsulation){
	var schema globalTypes.TableSchema = globalTypes.CreateTableSchema(tableName, tableSchema)
	AddTable(schema, allTableData)
}