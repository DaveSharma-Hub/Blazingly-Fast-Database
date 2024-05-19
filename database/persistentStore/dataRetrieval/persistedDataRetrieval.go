package persistedDataRetrieval

import (
	// "bufio"
    "fmt"
    // "log"
	"math"
	"bufio"
    "os"
	"io"
	"bytes"
	"strconv"
	"errors"
	"strings"
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/types"
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/persistentStore/binaryTree"
)


func check(e error, message string) {
    if e != nil {
		fmt.Println(message)
        panic(e)
    }
}

func CreateFile(fileName string, location string){
	finalFileLocation := location + fileName
	_, err := os.Create(finalFileLocation)
	check(err, "Error creating file")
}


func ReadFileFromLine(begByte int64, count int64, fileName string, innerKeyName string, innerKeyValue string, matchingOperator string, c chan string){
	// fmt.Println("CALLED")
	for i:=int64(0);i<count;i++{
		byteValue := begByte + (globalTypes.MAXPAYLOAD_BYTE_SIZE + 2)* i
		// fmt.Println("byte:"+strconv.Itoa(int(byteValue)))
		line,_ := GetPayloadByLineNumber(fileName, int64(byteValue))
		// fmt.Println("line:"+line)
		switch matchingOperator{
		case globalTypes.MATCHING_OPEQUAL:
			var reconstructed strings.Builder
			fmt.Fprintf(&reconstructed, "%s:{%s", innerKeyName, innerKeyValue)
			if strings.Contains(line, reconstructed.String()) {
				c <- line
			}
		default:
			//do nothing
		}
	}
	c <- "{}"
}

func GetLineNumber(filename string, key string)(int64,error){
	f,err := os.OpenFile(filename, os.O_APPEND|os.O_RDONLY|os.O_CREATE, 0600)
	if err != nil {
		check(err,"Error opening file for append")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lineNumber := -1
	for scanner.Scan(){
		substring := key+":"
		if strings.Contains(scanner.Text(),substring){
			item := strings.Split(scanner.Text(), ":")
			lineNumber,_ := strconv.Atoi(item[1])
			return int64(lineNumber),nil
		}
	}
	return int64(lineNumber), errors.New("NOT_FOUND")
}

func SearchLineConcurrently(fileName string, innerKeyName string, innerKeyValue string, matchingOperator string)(*globalTypes.Payload,error){
	var numberOfChannels int64 = 4 // get from number of processes by os maybe?
	
	var channels = make([]chan string,numberOfChannels)
	totalSize,_ := byteCounter(fileName)
	numberOfLines := int64(totalSize/ globalTypes.MAXPAYLOAD_BYTE_SIZE)
	if numberOfLines < numberOfChannels {
		numberOfChannels = 1
	}
	chunkSize := int64(math.Ceil(float64(numberOfLines/ numberOfChannels)))

	for i:=int64(0);i<numberOfChannels;i++{
		c := make(chan string)
		channels[i] = c
		beg := (i * chunkSize) * (globalTypes.MAXPAYLOAD_BYTE_SIZE)
		// fmt.Println("beg"+strconv.Itoa(int(beg)))
		// fmt.Println("chunkSize"+strconv.Itoa(int(chunkSize)))
		if beg + chunkSize*globalTypes.MAXPAYLOAD_BYTE_SIZE > totalSize {
			chunkSize = int64((totalSize - beg) / (globalTypes.MAXPAYLOAD_BYTE_SIZE))
		}
		go ReadFileFromLine(beg, chunkSize, fileName, innerKeyName, innerKeyValue, matchingOperator, c)
	}
	finalLine := "{}"
	for i:=int64(0);i<numberOfChannels;i++{
		output := <- channels[i]
		if output != "{}" {
			finalLine = output
		}
	}
	return globalTypes.ConvetBackToPayload(finalLine),nil
}
func GetPayloadByLineNumber(fileName string, lineNumber int64)(string,error){
	f,err := os.OpenFile(fileName, os.O_APPEND|os.O_RDONLY|os.O_CREATE, 0600)
	if err != nil {
		check(err,"Error opening file for append")
	}
	defer f.Close()

	_, err = f.Seek(lineNumber, 0)  // Set the current position for the fd
	if err != nil { // error handler
		return "{}",errors.New("NOT_FOUND")
	}

	reader := bufio.NewReader(f)
	line, _, err := reader.ReadLine()
	if err != nil {
		return "{}", errors.New("NOT_FOUND")
	}
	return string(line), nil
}

func lineCounter(fileName string) (int, error) {
	r, err := os.OpenFile(fileName, os.O_APPEND|os.O_RDONLY|os.O_CREATE, 0600)
	if err != nil {
		check(err,"Error opening file for append")
	}
    buf := make([]byte, 32*1024) // read 32K Bytes at a time and find '\n' separator
    count := 0
    lineSep := []byte{'\n'}

    for {
        c, err := r.Read(buf)
        count += bytes.Count(buf[:c], lineSep)
		fmt.Println(string(buf[:c]))
        switch {
        case err == io.EOF:
            return count, nil

        case err != nil:
            return count, err
        }
    }
	return count, nil
}

func byteCounter(fileName string)(int64,error){
	r, err := os.OpenFile(fileName, os.O_APPEND|os.O_RDONLY|os.O_CREATE, 0600)
	if err != nil {
		check(err,"Error opening file for append")
	}
	fi, err := r.Stat()
	if err != nil {
		// Could not obtain stat, handle error
		return 0, errors.New("Error finding byte size")
	}
	return fi.Size(),nil
}

func SetLineNumber(fileName string, fileNameMetaData string,key string)int64{
	byteNumber, _ := byteCounter(fileName)
	fmt.Println(byteNumber)
	f, err := os.OpenFile(fileNameMetaData, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		check(err,"Error opening file for append")
	}
	defer f.Close()

	var finalString strings.Builder
	fmt.Fprintf(&finalString,"%s:%d",key, byteNumber) 
	_, err = f.WriteString(finalString.String())
	check(err, "Faied to write to file")
	
	_, err = f.WriteString("\n")
	check(err, "Faied to write to file")

	return byteNumber
}

func SetPersistedDataFile(tableName string, key string, value *globalTypes.Payload)*binaryTree.DataMemoryLocation{
	var fileNameMetaData string = globalTypes.LOCATION + tableName + "_metaData.txt"
	var fileName string  = globalTypes.LOCATION + tableName + ".txt"

	var tablesDataFile string = globalTypes.LOCATION + "Tables.txt"
	tableFile, err := os.OpenFile(tablesDataFile, os.O_RDONLY, 0600)
	if err != nil {
		check(err,"Error opening file for append")
	}
	defer tableFile.Close()
	schema := [][]string{}

	scanner := bufio.NewScanner(tableFile)
	for scanner.Scan(){
		item := strings.Split(scanner.Text(), ":")
		if item[0] == tableName {
			if len(item)==3{
				keyTypeArr := strings.Split(item[1],")")
				for index := range(keyTypeArr){
					keyType := strings.Split(keyTypeArr[index], ",")
					if (len(keyType[0])>0 && len(keyType[1])>0) {
						key := keyType[0][1:len(keyType[0])]
						valueType := keyType[1]
						tmp := []string{key, valueType}
						schema = append(schema, tmp)
					}
				}
				break
			}
		}
	}

	isValid := globalTypes.VerifySchema(value, schema)
	if !isValid {
		// need to change what you return, can notify the user that incorrect schema
		fmt.Println("INCORRECT SCHEMA")
		return nil
	}

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		check(err,"Error opening file for append")
	}
	var byteNumber int64 = SetLineNumber(fileName, fileNameMetaData, key)

	defer file.Close()
	//	persist data in fileName, need to convert payload to string then convert back when getting

	var stringifiedPayload string = globalTypes.ConvertPayload(value)
	_,err = file.WriteString(stringifiedPayload)
	check(err, "Fail to write to file")
	
	_,err = file.WriteString("\n")
	check(err, "Fail to write to file")

	return &binaryTree.DataMemoryLocation{ByteOffset:byteNumber}
}

func AppendFileTableMeta(fileName string, location string, schema globalTypes.TableSchema) {
	finalLocation := location + fileName
	f, err := os.OpenFile(finalLocation, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		check(err,"Error opening file for append")
	}

	defer f.Close()
	tableName := schema.TableName + ":"

	_, err = f.WriteString(tableName)
	check(err, "Faied to write to file")

	for attributeKey := range schema.Attributes {
		dataId := attributeKey
		typeOfData := schema.Attributes[attributeKey].Type
		finalString  := "("+ dataId+","+ typeOfData +")"
		_, err := f.WriteString(finalString)
		check(err, "Faied to write to file")
	}
	var tableFileName string = schema.TableName + ".txt" 
	var tableFileNameMetaData string = schema.TableName + "_metaData.txt" 
	var tableFileNameString string = ": " + tableFileName

	_, err = f.WriteString(tableFileNameString)
	check(err, "Faied to write to file")

	CreateFile(tableFileName,globalTypes.LOCATION)
	CreateFile(tableFileNameMetaData,globalTypes.LOCATION)

	_, err = f.WriteString("\n")
	check(err, "Faied to write to file")

	// _,err := f.Close()
	// check(err, "Error closing file")
}


func GetPersistedDataFile(tableName string, key string, byteOffest int64)*globalTypes.Payload{
	fileName := globalTypes.LOCATION + tableName + ".txt"
	fileNameMetaData := globalTypes.LOCATION + tableName + "_metaData.txt"
	byteNumber := byteOffest
	if byteNumber == -1{
		bytes, err := GetLineNumber(fileNameMetaData, key)
		fmt.Println("bytes:"+strconv.Itoa(int(bytes)))
		byteNumber = bytes
		if err != nil{
			payload := globalTypes.CreateEmptyPayload()
			return &payload
			// ISSUE: Need to fix to return nil instead or a payload that indicates it doesnt exist
		}
	}

	stringifiedPayload, err := GetPayloadByLineNumber(fileName, byteNumber)
	if err != nil{
		payload := globalTypes.CreateEmptyPayload()
		return &payload
		// ISSUE: Need to fix to return nil instead or a payload that indicates it doesnt exist
	}
	return globalTypes.ConvetBackToPayload(stringifiedPayload)
}

func GetAllDataMatchingPersistedDataFile(tableName string, innerKeyName string, innerKeyValue string,matchingOperator string)*globalTypes.Payload{
	// this is where storage as a binary tree can help massively (maybe multi dimensional)
	// leverage more storage for smaller time complexity

	// fileName := globalTypes.LOCATION + tableName + ".txt"
	// fd, err := os.OpenFile(fileName, os.O_RDONLY, 0600)
	// if err!=nil {
	// 	check(err,"Error opening file for append")
	// }

	// defer fd.Close()
	// scanner := bufio.NewScanner(fd)
	
	// for scanner.Scan(){
	// 	text := scanner.Text()
	// 	switch matchingOperator{
	// 	case globalTypes.MATCHING_OPEQUAL:
	// 		var reconstructed strings.Builder
	// 		fmt.Fprintf(&reconstructed, "%s:{%s", innerKeyName, innerKeyValue)
	// 		if strings.Contains(text, reconstructed.String()) {
	// 			return globalTypes.ConvetBackToPayload(text)
	// 		}
	// 	default:
	// 		//do nothing
	// 	}
	// }

	// payload := globalTypes.CreateEmptyPayload()
	// return &payload

	// ====
	fileName := globalTypes.LOCATION + tableName + ".txt"

	payload, err := SearchLineConcurrently(fileName, innerKeyName, innerKeyValue, matchingOperator)

	if err != nil {
		panic("Error")
	}

	return payload
}

func matchString(superString string, subString string)(bool,int){
	var matching bool = false
	prevIndex := 0
	for i:=0;i<len(superString);i++{
		if superString[i] == subString[0]{
			matching = true
			prevIndex = i
			j:=0
			for ;j<len(subString) && i<len(superString);j++{
				if subString[j] == superString[i]{
					i++
				}else{
					matching = false
					i = prevIndex
					break
				}
			}
			if j<len(subString){
			    matching = false
			    i = prevIndex
			}
		}
		if matching{
			return true, prevIndex
		}
	}
	return false, -1
}

func getEndingByteRange(start int, line string)int{
	foundTwice := 0
	for i:=start;i<len(line);i++{
		if foundTwice==2 && line[i] == '}'{
			return i
		}
		if foundTwice<2 && line[i] == ':'{
			foundTwice++
		}
	}
	return -1
}

func UpdatePersistedDataFile(tableName string, key string, byteOffset int64, payload *globalTypes.Payload){
	fmt.Println("CALLED")
	fileNameMetaData := globalTypes.LOCATION + tableName + "_metaData.txt"
	fileName := globalTypes.LOCATION + tableName + ".txt"
	bytes := byteOffset
	if bytes == -1{
		bytesValue, err := GetLineNumber(fileNameMetaData, key)
		bytes = bytesValue
		if err != nil {
			fmt.Println("ERROR")
			return
		}
	}
	line,err := GetPayloadByLineNumber(fileName, bytes)
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	f, err := os.OpenFile(fileName, os.O_CREATE | os.O_WRONLY, os.ModeAppend)
    if err != nil {
        panic("File not found")
    }
	defer f.Close()

	/// TODO rename this function to replace data, and create update to update entire data with new payload
	originalPayload := globalTypes.ConvetBackToPayload(line)
	
	for keys := range(payload.Item){
		_,ok := originalPayload.Item[keys]
		if ok{
			originalPayload.Item[keys] = payload.Item[keys]
		}
	}
	
	globalTypes.FillPayloadTillMax(originalPayload)
	newPayloadAsAString := globalTypes.ConvertPayload(originalPayload)
	whence := io.SeekStart
	_, err = f.Seek(int64(bytes), whence)
	f.WriteString(newPayloadAsAString)
	f.Sync() //flush to disk
}

func RemoveDataPersistedFile(tableName string, key string, byteOffset int64){
	fileNameMetaData := globalTypes.LOCATION + tableName + "_metaData.txt"
	fileName := globalTypes.LOCATION + tableName + ".txt"
	bytes := byteOffset
	if bytes == -1 {
		bytesValue, err := GetLineNumber(fileNameMetaData, key)
		bytes = bytesValue
		if err != nil {
			fmt.Println("ERROR")
			return
		}
	}
	f, err := os.OpenFile(fileName, os.O_CREATE | os.O_WRONLY, os.ModeAppend)
    if err != nil {
        panic("File not found")
    }
	defer f.Close()

	whence := io.SeekStart
	_, err = f.Seek(bytes, whence)
	if err != nil {
        panic("Byte range not found")
    }

	var replaceData strings.Builder
	fmt.Fprintf(&replaceData, "%s%s%s", "{",globalTypes.EMPTY_KEY,":{")
	for i:=0;i<globalTypes.MAXPAYLOAD_BYTE_SIZE-globalTypes.RemovedSize-2;i++{
		fmt.Fprintf(&replaceData, "%s", globalTypes.EMPTY_VALUE)
	}	
	fmt.Fprintf(&replaceData, "%s", ":string},}\n")

	f.WriteString(replaceData.String())
	f.Sync() //flush to disk
}