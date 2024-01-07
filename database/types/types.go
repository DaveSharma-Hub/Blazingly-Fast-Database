package globalTypes

type CommandLineArguments struct {
	CacheMaxSize int64
}
 
type AtomicItem struct{
	Value string `json:"value"`
	Type string `json:"type"`
}

// type databaseOutput struct{
// 	Output string `json:"output"`
// 	Data string `json:"data"`
// }

type Payload struct{
    Item map[string] AtomicItem
}

func CreatePayload(input [][]string)Payload{
	
	var newItem Payload
	newItem.Item = make(map[string] AtomicItem)

	for i:= 0; i<len(input); i++{
		item := input[i]
		key := item[0]
		value := item[1]
		typeOfValue := item[2]
		newItem.Item[key] = AtomicItem{Value:value, Type:typeOfValue}
	}
	return newItem
}

func CreateEmptyPayload()Payload{
	var newItem Payload
	newItem.Item = make(map[string] AtomicItem)
	return newItem
}
