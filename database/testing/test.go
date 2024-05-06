package testing

import (
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/persistentStore/dataRetrieval"
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/persistentStore"
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/types"
	"fmt"
)

func MainTest(){
	perisistentStore := persistentStoreClient.InitPersistentStoreClient();
	value := globalTypes.CreatePayload([][]string{{"id","1","string"},{"name","John","string"},{"age","24","integer"},{"occupation","Engineer","string"}})
	value2 := globalTypes.CreatePayload([][]string{{"id","2","string"},{"name","Bob","string"},{"age","65","integer"},{"occupation","Carpenter","string"}})
	value3 := globalTypes.CreatePayload([][]string{{"name","Bobjnkjnjk","string"},{"age","15","integer"},{"occupation","Manager of Architect","string"}})
	// persistedDataRetrieval.SetPersistedDataFile("Users", "1", &value)

	persistentStoreClient.SetData("Users","1",value, perisistentStore)
	persistentStoreClient.SetData("Users","2",value2, perisistentStore)
	persistentStoreClient.UpdateData("Users","1",value3, perisistentStore)
	// // str:= globalTypes.ConvertPayload(&value)
	// // va := globalTypes.ConvetBackToPayload(str)
	// // fmt.Println(globalTypes.ConvertPayload(va))
	// // v := persistentStoreClient.GetMatchingData("Users","name", "Bob", "EQUAL", perisistentStore)
	
	persistedDataRetrieval.RemoveDataPersistedFile("Users", "1", -1)
	persistedDataRetrieval.RemoveDataPersistedFile("Users", "1", -1)
	v:= persistentStoreClient.GetData("Users","1", perisistentStore)
	str := globalTypes.ConvertPayload(&v)
	fmt.Println(str)

	// payload,_ := globalTypes.FillPayloadTillMax(&value)

	// str := globalTypes.ConvertPayload(payload)
	// load := globalTypes.ConvetBackToPayload(str)
	// strAgain := globalTypes.ConvertPayload(load)
	// fmt.Println(strAgain)

}
