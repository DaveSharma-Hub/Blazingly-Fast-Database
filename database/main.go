package main

import (
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/commandArgs"
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/server"
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/cache"
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/persistentStore"
	// "github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/types"
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/dataCacheClient"
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/testing"
	"os"
) 

func main(){
		//   node1 := initNode("1",payload{value:"1"})
		//   node2 := initNode("2",payload{value:"2"})
		//   node3 := initNode("3", payload{value:"3"})
		//   node4 := initNode("4", payload{value:"4"})
		//   node5 := initNode("5",payload{value:"5"})
		//   node6 := initNode("6",payload{value:"6"})
		//   LL := initLinkedList();
		//   pushToLinkedList(LL,node1)
		//   pushToLinkedList(LL,node2)
		//   pushToLinkedList(LL,node3)
		//   pushToLinkedList(LL,node4)
		//   pushToLinkedList(LL,node5)
		//   pushToLinkedList(LL,node6)
		//   printLL(LL)
		// lru := cacheClient.InitCacheClient(3)
		// cacheClient.ExecuteOperation(lru,"1",func() string{
		// 	return "1"
		// })
		// cacheClient.ExecuteOperation(lru,"2", func() string{
		// 	return "2"
		// })
		// cacheClient.ExecuteOperation(lru,"3", func() string{
		// 	return "3"
		// })
		// cacheClient.ExecuteOperation(lru,"4", func() string{
		// 	return "4"
		// })
		// cacheClient.ExecuteOperation(lru,"5", func() string{
		// 		return "5"
		// 	})
		// cacheClient.ExecuteOperation(lru,"6", func() string{
		// 		return "6"
		// 	})
		// cacheClient.ExecuteOperation(lru,"1", func() string{
		// 		return "1"
		// 	})
		// cacheClient.ExecuteOperation(lru,"2", func() string{
		// 		return "2"
		// 	})
		// cacheClient.ExecuteOperation(lru,"1", func() string{
		// 		return "1"
		// })
		// // // printLL(lru.list)	
		// cacheClient.PrintLL(lru)

	
	inputArguments := os.Args[1:]

    arguments := commandArgs.ParseInput(inputArguments)

	if arguments.IsTesting {
		testing.MainTest()
		return
	}

	perisistentStore := persistentStoreClient.InitPersistentStoreClient();
	dataQueryInMemoryCacheClient := cacheClient.InitCacheClient(arguments.CacheMaxSize);


	dataAndCacheClient := dataCacheClient.CreateFunctionMapWrapper(dataQueryInMemoryCacheClient, perisistentStore)

	// cacheClient.ExecuteOperation(dataQueryInMemoryCacheClient,"1", perisitentStoreClient.GetData)

	// cacheClient.ExecuteOperation(dataQueryInMemoryCacheClient,"1",func(key string) globalTypes.Payload{
	// 	return globalTypes.CreatePayload([][]string{{"id","1","string"},{"name","John","string"}})
	// })

	clientRouter := server.InitServer(dataAndCacheClient)
	server.RunServer(clientRouter)


	// pass in the cache and perisistent store into server
	// can parse input string and return either from cache or persistent store

}