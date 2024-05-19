package testing

import (
	// "github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/persistentStore/dataRetrieval"
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/persistentStore"
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/types"
	"fmt"
	"strconv"
	"time"
)

/*
schema -> {"keyName","type"}
*/

func createFakeDate(schema [][]string, count int)[]globalTypes.Payload{
	var fake []globalTypes.Payload = make([]globalTypes.Payload, count)
	for i:=0;i<count;i++{
		var data [][]string = make([][]string, len(schema))
		for j:=0;j<len(schema);j++{
			var inner []string = make([]string, len(schema[j]))
			for z:=0;z<len(schema[j]);z++{
				value := ""
				if z==0 || z==2{
					value = schema[j][z]
				}else{
					value = "User"+strconv.Itoa(i)
				}
				inner[z] = value

			}
			data[j] = inner
		}
		fake[i] = globalTypes.CreatePayload(data)
	}
	return fake
}


func createAndAssignFakeData(schema [][]string, count int, perisistentStore *persistentStoreClient.TableEncapsulation){
	var payloads = createFakeDate(schema, count)
	fmt.Println("COUNT:"+strconv.Itoa(len(payloads)))
	for i:=0;i<len(payloads);i++{
		// fmt.Println(payloads[i].Item["age"].Type)
		persistentStoreClient.SetData("Users",strconv.Itoa(i),payloads[i], perisistentStore)
	}

}

func MainTest(){
	perisistentStore := persistentStoreClient.InitPersistentStoreClient();
	// value := globalTypes.CreatePayload([][]string{{"id","1","string"},{"name","John","string"},{"age","24","integer"},{"occupation","Engineer","string"}})
	// value2 := globalTypes.CreatePayload([][]string{{"id","2","string"},{"name","Bob","string"},{"age","65","integer"},{"occupation","Carpenter","string"}})
	// value3 := globalTypes.CreatePayload([][]string{{"name","Bobjnkjnjk","string"},{"age","15","integer"},{"occupation","Manager of Architect","string"}})
	// value4 := globalTypes.CreatePayload([][]string{{"name","Bobjnkjnjk","string"},{"age","15","integer"},{"occupation","Manager of Architect","string"}})
	// value5 := globalTypes.CreatePayload([][]string{{"name","Bobjnkjnjk","string"},{"age","15","integer"},{"occupation","Manager of Architect","string"}})
	
	count := 1000
	var schema = [][]string{{"id","2","string"},{"name","Bobjnkjnjk","string"},{"age","15","integer"},{"occupation","Manager of Architect","string"}}
	createAndAssignFakeData(schema, count, perisistentStore)
	// persistedDataRetrieval.SetPersistedDataFile("Users", "1", &value)

	// persistentStoreClient.SetData("Users","1",value, perisistentStore)
	// persistentStoreClient.SetData("Users","2",value2, perisistentStore)
	// persistentStoreClient.SetData("Users","3",value3, perisistentStore)
	// persistentStoreClient.UpdateData("Users","1",value3, perisistentStore)
	// // str:= globalTypes.ConvertPayload(&value)
	// // va := globalTypes.ConvetBackToPayload(str)
	// // fmt.Println(globalTypes.ConvertPayload(va))
	// v := persistentStoreClient.GetMatchingData("Users","name", "Bob", "EQUAL", perisistentStore)
	
	// persistedDataRetrieval.RemoveDataPersistedFile("Users", "1", -1)
	// persistedDataRetrieval.RemoveDataPersistedFile("Users", "1", -1)
	beg := time.Now().UnixNano()
	v:= persistentStoreClient.GetData("Users",strconv.Itoa(count-2), perisistentStore)
	str := globalTypes.ConvertPayload(&v)
	fmt.Println(str)

	end := time.Now().UnixNano()
	fmt.Println("Diff:"+strconv.Itoa(int(end-beg)))

	// payload,_ := globalTypes.FillPayloadTillMax(&value)

	// str := globalTypes.ConvertPayload(payload)
	// load := globalTypes.ConvetBackToPayload(str)
	// strAgain := globalTypes.ConvertPayload(load)
	// fmt.Println(strAgain)

}
