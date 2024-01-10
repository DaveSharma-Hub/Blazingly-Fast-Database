package globalTypes

type CommandLineArguments struct {
	CacheMaxSize int64
}
 
type AtomicItem struct{
	Value string `json:"value"`
	Type string `json:"type"`
}

type AtomicAttribute struct{
	Type string `json:"type"`
}

type Payload struct{
    Item map[string] AtomicItem `json:"item"`
}

type TableSchema struct{
	Attributes map[string] AtomicAttribute
	TableName string
}

func CreateTableSchema(tableName string, input [][]string)TableSchema{
	var newSchema TableSchema
	newSchema.Attributes = make(map[string] AtomicAttribute)
	for i:=0; i<len(input); i++ {
		attribute := input[i]
		name := attribute[0]
		attributeType := attribute[1]
		newSchema.Attributes[name] = AtomicAttribute{Type:attributeType}
	}
	newSchema.TableName = tableName
	return newSchema
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
