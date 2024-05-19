package globalTypes

import (
	"strings"
	"fmt"
	"strconv"
	"errors"
)

const LOCATION = "./rawData/"
const MATCHING_OPEQUAL = "EQUAL"
const MAXPAYLOAD_BYTE_SIZE = 2500
const EMPTY_KEY = "EMPTY_KEY"
const EMPTY_VALUE = "+"
// Maybe change to multiple of 2 instead
// const RemovedSize = 5+9+6+1 // size of colons and brackets + EMPTY_KEY + string size
const RemovedSize = 9+6+2+2 // size of colons and brackets + EMPTY_KEY + string size
// 9 = length of "EMPTY_KEY", 6 = length of "string", 2 = # of colons, 2 = # of parenthesis.

type OtherClientPassedInfo struct{
	InnerKey string
	InnerKeyValue string
	Comparator string
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
		if key != EMPTY_KEY{
			value := payload.Item[key]
			line := key + ":{" + value.Value + ":" + value.Type + "},"
			fmt.Fprintf(&finalStr, "%s", line)
		}
	}
	lastLine, ok := payload.Item[EMPTY_KEY]

	if ok{
		value := lastLine
		line := EMPTY_KEY + ":{" + value.Value + ":" + value.Type + "},"
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
			if key != EMPTY_KEY{
				value := items[1][1:]
				valueType := items[2][0:len(items[2])-1]
				tmp := []string{key,value,valueType}
				input = append(input, tmp)
			}
		}
		fmt.Println("keyValuePayload", keyValuePayload)
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
				fmt.Println("VALUE:"+valueType+" - "+value.Type)
				return false
			}
		}else{
			fmt.Println("KEY"+key)
			return false
		}
	}
	return true	
}

func getPayloadSize(payload *Payload)uint64{
	var convertedPayload string = ConvertPayload(payload)
	return (uint64)(len(convertedPayload)) 
}

func FillPayloadTillMax(payload *Payload)(*Payload,error){
	currentSize := getPayloadSize(payload)
	fmt.Println("currentSize", currentSize)
	fmt.Println("MAXPAYLOAD_BYTE_SIZE", MAXPAYLOAD_BYTE_SIZE)
	fmt.Println("first", currentSize < MAXPAYLOAD_BYTE_SIZE)
	fmt.Println("second", (currentSize + RemovedSize) < MAXPAYLOAD_BYTE_SIZE)

	if currentSize < MAXPAYLOAD_BYTE_SIZE && (currentSize + RemovedSize < MAXPAYLOAD_BYTE_SIZE){
		extraNeeded := MAXPAYLOAD_BYTE_SIZE - (currentSize + RemovedSize)
		var addedValue strings.Builder
		for i:= (uint64)(0);i<extraNeeded;i++{
			fmt.Fprintf(&addedValue, "%s", EMPTY_VALUE)
		}
		payload.Item[EMPTY_KEY] = AtomicItem{Value:addedValue.String(), Type:"string"}
		return payload, nil
	}else{
		strSize:= strconv.Itoa(MAXPAYLOAD_BYTE_SIZE)
		errorMessage := "Invalid payload size: greater than" +  strSize
		return payload, errors.New(errorMessage)
	}
}