package globalTypes

type CommandLineArguments struct {
	CacheMaxSize int64
}
 
type AtomicItem struct{
	value string `json:"value"`
	typeOfValue string `json:"typeOfValue"`
}

// type databaseOutput struct{
// 	Output string `json:"output"`
// 	Data string `json:"data"`
// }

type Payload struct{
    item map[string] AtomicItem
}

func CreatePayload(input [][]string)Payload{
	
	var newItem Payload
	newItem.item = make(map[string] AtomicItem)

	for i:= 0; i<len(input); i++{
		item := input[i]
		key := item[0]
		value := item[1]
		typeOfValue := item[2]
		newItem.item[key] = AtomicItem{value:value, typeOfValue:typeOfValue}
	}
	return newItem
}

func CreateEmptyPayload()Payload{
	var newItem Payload
	newItem.item = make(map[string] AtomicItem)
	return newItem
}
