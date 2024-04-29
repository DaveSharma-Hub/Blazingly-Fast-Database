package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/types"
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/dataCacheClient"
)

type FunctionWrapperType func(*gin.Context, dataCacheClient.DataCacheExecutionType)

type PostQueryInputType struct{
	TableName string `json:"table_name" binding:"required"`
	PartitionKey string `json:"partition_key" binding:"required"`
}

type ExtendedPayloadStructure struct{
	Value string `json:"value" binding:"required"`
	ExtendedValue [][]ExtendedPayloadStructure `json:"payload" binding:"required"`
}

type PostSetDataInputType struct{
	TableName string `json:"table_name" binding:"required"`
	PartitionKey string `json:"partition_key" binding:"required"`
	DataPayload  [][]string `json:"payload" binding:"required"`
}

func GetPayloadFromPostSetDataInput(input PostSetDataInputType)globalTypes.Payload{
	return globalTypes.CreatePayload(input.DataPayload)
}