package testing

import (
	// "github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/persistentStore/dataRetrieval"
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/persistentStore"
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/types"
	"fmt"
)

func MainTest(){
	perisistentStore := persistentStoreClient.InitPersistentStoreClient();
	value := globalTypes.CreatePayload([][]string{{"id","1","string"},{"name","John","string"},{"age","24","integer"},{"occupation","Engineer","string"}})
	value2 := globalTypes.CreatePayload([][]string{{"id","2","string"},{"name","Bob","string"},{"age","65","integer"},{"occupation","Carpenter","string"}})
	// persistedDataRetrieval.SetPersistedDataFile("Users", "1", &value)

	persistentStoreClient.SetData("Users","1",value, perisistentStore)
	persistentStoreClient.SetData("Users","2",value2, perisistentStore)
	// str:= globalTypes.ConvertPayload(&value)
	// va := globalTypes.ConvetBackToPayload(str)
	// fmt.Println(globalTypes.ConvertPayload(va))
	v := persistentStoreClient.GetData("Users","2", perisistentStore)
	str := globalTypes.ConvertPayload(&v)
	fmt.Println(str)
}
