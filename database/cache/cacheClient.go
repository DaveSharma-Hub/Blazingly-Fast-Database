package cacheClient

import (
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/cache/lruCache"
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/types"
)

func InitCacheClient(maxCachedItems int64)*lruCache.LRUCache{
	inMemoryLRUCache := lruCache.InitLRUCache(maxCachedItems);
	return inMemoryLRUCache
}

func ExecuteOperationGetItem(inMemoryLRUCache *lruCache.LRUCache, key string, fnCallback lruCache.PersistedItemConversion)globalTypes.Payload{
	return lruCache.GetItem(inMemoryLRUCache, key, fnCallback)
}

func ExecuteOperationSetItem(inMemoryLRUCache *lruCache.LRUCache, key string, dataPayload globalTypes.Payload){
	lruCache.SetItem(inMemoryLRUCache, key, dataPayload)
}

func ExecuteOperationUpdateItem(inMemoryLRUCache *lruCache.LRUCache, key string, dataPayload globalTypes.Payload){
	lruCache.UpdateItem(inMemoryLRUCache, key, dataPayload)
}

func ExecuteOperationRemoveItem(inMemoryLRUCache *lruCache.LRUCache, key string){
	lruCache.RemoveItem(inMemoryLRUCache, key)
}

func PrintLL(inMemoryLRUCache *lruCache.LRUCache){
	lruCache.PrintItems(inMemoryLRUCache)
}	

