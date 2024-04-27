package persistedDataRetrieval

import (
	// "bufio"
    "fmt"
    // "log"
    "os"
	"io"
	"bytes"
	"strconv"
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


func GetLineNumber(filename string, key string)int{
	return 0
}

func lineCounter(r io.Reader) (int, error) {
    buf := make([]byte, 32*1024) // read 32K Bytes at a time and find '\n' separator
    count := 0
    lineSep := []byte{'\n'}

    for {
        c, err := r.Read(buf)
        count += bytes.Count(buf[:c], lineSep)

        switch {
        case err == io.EOF:
            return count, nil

        case err != nil:
            return count, err
        }
    }
	return count, nil
}

func SetLineNumber(file io.Reader, fileNameMetaData string,key string)int{
	lineNumber, _ := lineCounter(file)

	f, err := os.OpenFile(fileNameMetaData, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		check(err,"Error opening file for append")
	}
	defer f.Close()

	finalString := key +":" + strconv.Itoa(lineNumber)
	_, err = f.WriteString(finalString)
	check(err, "Faied to write to file")

	return lineNumber
}

func SetPersistedDataFile(tableName string, key string, value *globalTypes.Payload)*binaryTree.DataMemoryLocation{
	var fileNameMetaData string = globalTypes.LOCATION + tableName + "_metaData.txt"
	var fileName string  = globalTypes.LOCATION + tableName + ".txt"

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_RDONLY| os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		check(err,"Error opening file for append")
	}
	var lineNumber int = SetLineNumber(file, fileNameMetaData, key)

	defer file.Close()
	//	persist data in fileName, need to convert payload to string then convert back when getting
	var stringifiedPayload string = globalTypes.ConvertPayload(value)
	_,err = file.WriteString(stringifiedPayload)
	check(err, "Fail to write to file")

	return &binaryTree.DataMemoryLocation{LineNumber:lineNumber}
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


