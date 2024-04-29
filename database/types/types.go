package globalTypes

import (
	"strings"
	"fmt"
)

const LOCATION = "./rawData/"
const MATCHING_OPEQUAL = "EQUAL"

type OtherClientPassedInfo struct{
	InnerKey string
	InnerKeyValue string
}

type CommandLineArguments struct {
	CacheMaxSize int64
	IsTesting bool
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

func ConvertPayload(payload *Payload)string{
	var finalStr strings.Builder
	fmt.Fprintf(&finalStr, "{")
	
	for key := range(payload.Item){
		value := payload.Item[key]
		line := key + ":{" + value.Value + ":" + value.Type + "},"
		fmt.Fprintf(&finalStr, "%s", line)
	}
	fmt.Fprintf(&finalStr, "}")
	return finalStr.String()
}
//ISSUE: What if comma, or {} are in value then how to handle parsing

func ConvetBackToPayload(payload string)*Payload{
    input := [][]string{}

	strippedPayload := payload[1:len(payload)-1]
	splitPayload := strings.Split(strippedPayload, ",")

	for index := range(splitPayload){
	    keyValuePayload := splitPayload[index]
		items := strings.Split(keyValuePayload, ":")
		if len(items)==3{
    		key := items[0]
    		value := items[1][1:]
    		valueType := items[2][0:len(items[2])-1]
    		tmp := []string{key,value,valueType}
    		input = append(input, tmp)
		}
	}

	newPayload := CreatePayload(input)
	return &newPayload
}
//ISSUE: What if comma, or {} are in value then how to handle parsing


func VerifySchema(payload *Payload, schema [][]string)bool{
	for index := range(schema){
		keyValue := schema[index]
		key := keyValue[0]
		valueType := keyValue[1]
		value, ok := payload.Item[key]
		if ok{
			if valueType != value.Type{
				return false
			}
		}else{
			return false
		}
	}
	return true	
}