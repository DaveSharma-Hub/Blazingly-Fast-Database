package persistentStoreClient

import (
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/types"
	// "github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/cache/binaryTree"
	"time"
)

func tempStoreData(){

}

func InitPersistentStoreClient(){
	tempStoreData()

}

func GetData(key string) globalTypes.Payload{
	time.Sleep(1000000000)
	return globalTypes.CreatePayload([][]string{{"id","1","string"},{"name","John","string"},{"age","14","integer"},{"occupation","Engineer","string"}})
}

func SetData(key string, value globalTypes.Payload){

}