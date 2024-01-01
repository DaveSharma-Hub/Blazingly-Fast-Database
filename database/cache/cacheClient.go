package cacheClient

import (
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/cache/lruCache"
)

func InitCacheClient(maxCachedItems int64)*lruCache.LRUCache{
	inMemoryLRUCache := lruCache.InitLRUCache(maxCachedItems);
	return inMemoryLRUCache
}

func ExecuteOperation(inMemoryLRUCache *lruCache.LRUCache, key string, fnCallback lruCache.PersistedItemConversion)lruCache.Payload{
	return lruCache.GetItem(inMemoryLRUCache, key, fnCallback);
}

func PrintLL(inMemoryLRUCache *lruCache.LRUCache){
	lruCache.PrintItems(inMemoryLRUCache)
}	

