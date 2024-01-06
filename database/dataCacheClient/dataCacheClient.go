package dataCacheClient

import (
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/cache"
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/persistentStore"
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/types"
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/cache/lruCache"
)

type DataCacheExecutionType func(string, globalTypes.Payload)globalTypes.Payload

type DataCacheClientReturnType map[string] DataCacheExecutionType


func CreateFunctionMapWrapper(dataQueryInMemoryCacheClient *lruCache.LRUCache)DataCacheClientReturnType{
	functionMap := make(map[string] DataCacheExecutionType)

	functionMap["GetData"] = func(key string, value globalTypes.Payload)globalTypes.Payload{
			return cacheClient.ExecuteOperation(dataQueryInMemoryCacheClient,key, func(storeKey string)globalTypes.Payload{
				return persistentStoreClient.GetData(storeKey)
			})
		}

	return functionMap
}