package persistentStoreClient

import (
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/types"
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/persistentStore/binaryTree"
	// "time"
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/persistentStore/dataRetrieval"
	"fmt"
)

type TableMetaDataType struct{
	TableAttributes globalTypes.TableSchema
}



type DataReturnType struct{
	Payload *globalTypes.Payload
	DataLocation *binaryTree.DataMemoryLocation
}

// here is where you can optimize using binary tree's etc
type TableDataType struct{
	Data map[string] *DataReturnType	// Array of payload that matches schema
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
	tableData := TableDataType{Data:make(map[string] *DataReturnType)}
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

	// SetData("Users","First",globalTypes.CreatePayload([][]string{{"id","1","string"},{"name","John","string"},{"age","24","integer"},{"occupation","Engineer","string"}}),allTableData)
	// SetData("Users","Second",globalTypes.CreatePayload([][]string{{"id","2","string"},{"name","Bob","string"},{"age","74","integer"},{"occupation","Plumber","string"}}),allTableData)
	// SetData("Users","Third",globalTypes.CreatePayload([][]string{{"id","3","string"},{"name","Kelly","string"},{"age","44","integer"},{"occupation","Financy","string"}}),allTableData)

	
	
	schema := globalTypes.CreateTableSchema("Users",[][]string{{"id","string"},{"name","string"},{"age","integer"},{"occupation","string"}})
	schema2 := globalTypes.CreateTableSchema("Locations",[][]string{{"id","string"},{"city","string"},{"country","string"}})

	persistedDataRetrieval.CreateFile("Tables.txt", globalTypes.LOCATION)
	persistedDataRetrieval.AppendFileTableMeta("Tables.txt",globalTypes.LOCATION,schema)
	persistedDataRetrieval.AppendFileTableMeta("Tables.txt",globalTypes.LOCATION,schema2)
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
					var byteOffset int64 = -1
					//  if allTableData.TableInformation[tableName].TableData.Data[key].DataLocation != nil{
					// 	byteOffset = allTableData.TableInformation[tableName].TableData.Data[key].DataLocation.ByteOffset
					// }
					return *persistedDataRetrieval.GetPersistedDataFile(tableName, key, byteOffset)
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
				if allTableData.TableInformation[tableName].TableData.Data[key] == nil {
					payload, err := globalTypes.FillPayloadTillMax(&value)
					if err != nil {
						// need to return an error to client
						fmt.Println("ERROR")
					}else{
						fmt.Println("WRITING")

						var dataMemLocation *binaryTree.DataMemoryLocation = persistedDataRetrieval.SetPersistedDataFile(tableName, key, payload)
						allTableData.TableInformation[tableName].TableData.Data[key] = &DataReturnType{Payload:nil,DataLocation:dataMemLocation}
					}
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
				if allTableData.TableInformation[tableName].TableData.Data[key] != nil {
					_, okAgain := allTableData.TableInformation[tableName].TableData.Data[key]
					if okAgain {
						var byteOffset int64 = -1
						if allTableData.TableInformation[tableName].TableData.Data[key].DataLocation != nil{
							byteOffset = allTableData.TableInformation[tableName].TableData.Data[key].DataLocation.ByteOffset
						}
						persistedDataRetrieval.UpdatePersistedDataFile(tableName, key, byteOffset, &value)

					}else{
						SetData(tableName,key, value, allTableData)
					}
				}
			}
		}
	}
}

func RemoveData(tableName string, key string, value globalTypes.Payload, allTableData *TableEncapsulation){
	_, ok := allTableData.TableInformation[tableName]
	if ok {
		if allTableData.TableInformation[tableName].TableData != nil {
			if allTableData.TableInformation[tableName].TableData.Data != nil {
				if allTableData.TableInformation[tableName].TableData.Data[key] != nil {
					_, okAgain := allTableData.TableInformation[tableName].TableData.Data[key]
					if okAgain {
						var byteOffset int64 = -1
						if allTableData.TableInformation[tableName].TableData.Data[key].DataLocation != nil{
							byteOffset = allTableData.TableInformation[tableName].TableData.Data[key].DataLocation.ByteOffset
						}
						persistedDataRetrieval.RemoveDataPersistedFile(tableName, key, byteOffset)
						delete(allTableData.TableInformation[tableName].TableData.Data,key)
					}
				}
			}
		}
	}
}

func GetMatchingData(tableName string, innerKey string, innerMatchingValue string, comparator string,allTableData *TableEncapsulation)globalTypes.Payload{
	_, ok := allTableData.TableInformation[tableName]
	if ok {
		if allTableData.TableInformation[tableName].TableData != nil {
			if allTableData.TableInformation[tableName].TableData.Data != nil {
				return *persistedDataRetrieval.GetAllDataMatchingPersistedDataFile(tableName, innerKey, innerMatchingValue, comparator)
			}
		}
	}
	return globalTypes.CreateEmptyPayload();
}

func CreateTable(tableName string, tableSchema[][]string, allTableData *TableEncapsulation){
	var schema globalTypes.TableSchema = globalTypes.CreateTableSchema(tableName, tableSchema)
	AddTable(schema, allTableData)
}