package testing

import (
	// "github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/persistentStore/dataRetrieval"
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/persistentStore"
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/types"
	// "fmt"
)

func MainTest(){
	perisistentStore := persistentStoreClient.InitPersistentStoreClient();
	value := globalTypes.CreatePayload([][]string{{"id","1","string"},{"name","John","string"},{"age","24","integer"},{"occupation","Engineer","string"}})
	// persistedDataRetrieval.SetPersistedDataFile("Users", "1", &value)

	persistentStoreClient.SetData("Users","1",value, perisistentStore)
	// str:= globalTypes.ConvertPayload(&value)
	// va := globalTypes.ConvetBackToPayload(str)
	// fmt.Println(globalTypes.ConvertPayload(va))
}
