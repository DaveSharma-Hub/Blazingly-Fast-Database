package persistentStoreClient

import (
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/types"
	// "github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/cache/binaryTree"
)

func tempStoreData(){

}

func InitPersistentStoreClient(){
	tempStoreData()

}

func GetData(key string) globalTypes.Payload{
	return globalTypes.CreatePayload([][]string{{"id","1","string"},{"name","John","string"}})
}

func SetData(key string, value globalTypes.Payload){

}