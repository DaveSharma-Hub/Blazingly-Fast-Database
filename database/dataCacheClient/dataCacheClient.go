package dataCacheClient

import (
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/cache"
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/persistentStore"
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/types"
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/cache/lruCache"
)

type DataCacheExecutionType func(string,string, globalTypes.Payload, *globalTypes.OtherClientPassedInfo)globalTypes.Payload

type DataCacheClientReturnType map[string] DataCacheExecutionType


func CreateFunctionMapWrapper(dataQueryInMemoryCacheClient *lruCache.LRUCache, dataStore *persistentStoreClient.TableEncapsulation)DataCacheClientReturnType{
	functionMap := make(map[string] DataCacheExecutionType)

	functionMap["QueryData"] = func(tableName string, key string, value globalTypes.Payload, otherInfo *globalTypes.OtherClientPassedInfo)globalTypes.Payload{
		return cacheClient.ExecuteOperationGetItem(dataQueryInMemoryCacheClient,key, func(storeKey string)globalTypes.Payload{
			return persistentStoreClient.GetData(tableName, storeKey, dataStore)
		})
	}
	functionMap["SetData"] = func(tableName string, key string, value globalTypes.Payload, otherInfo *globalTypes.OtherClientPassedInfo)globalTypes.Payload{
		persistentStoreClient.SetData(tableName, key, value, dataStore)
		cacheClient.ExecuteOperationSetItem(dataQueryInMemoryCacheClient,key, value)
		//since returning empty globalType.Payload <- can find a better way later on
		return globalTypes.CreateEmptyPayload()
	}

	functionMap["UpdateData"] = func(tableName string, key string, value globalTypes.Payload, otherInfo *globalTypes.OtherClientPassedInfo)globalTypes.Payload{
		persistentStoreClient.UpdateData(tableName, key, value, dataStore)
		cacheClient.ExecuteOperationUpdateItem(dataQueryInMemoryCacheClient,key, value)
		//since returning globalType.Payload <- can find a better way later on
		return globalTypes.CreateEmptyPayload()
	}

	functionMap["QueryMatchingData"] = func(tableName string, key string, value globalTypes.Payload, otherInfo *globalTypes.OtherClientPassedInfo)globalTypes.Payload{
		return persistentStoreClient.GetMatchingData(tableName, otherInfo.InnerKey, otherInfo.InnerKeyValue, otherInfo.Comparator, dataStore)
	}
	
	functionMap["RemoveData"] = func(tableName string, key string, value globalTypes.Payload, otherInfo *globalTypes.OtherClientPassedInfo)globalTypes.Payload{
		cacheClient.ExecuteOperationRemoveItem(dataQueryInMemoryCacheClient, key)
		persistentStoreClient.RemoveData(tableName, key, value, dataStore)
		return globalTypes.CreateEmptyPayload()
	}

	return functionMap
}