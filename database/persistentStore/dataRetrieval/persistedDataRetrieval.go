package persistedDataRetrieval

import (
	// "bufio"
    "fmt"
    // "log"
    "os"
	"io"
	"bytes"
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
	finalLocation := location+"/"+fileName
	_, err := os.Create(finalLocation)
	check(err, "Error creating file")
}


func GetLineNumber(filename string, key string)int{
	return 0
}

func lineCounter(r io.Reader) (int, error) {
    buf := make([]byte, 32*1024)
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
}

func SetLineNumber(filename string, key string)int{
	f, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0600)
	if err != nil {
		check(err,"Error opening file for append")
	}
	defer f.Close()
	lineNumber, _ := lineCounter(f)

	finalString := key +":"+ string(lineNumber)
	_, err = f.WriteString(finalString)
	check(err, "Faied to write to file")

	return lineNumber
}

func SetPersistedDataFile(tableName string, key string, value *globalTypes.Payload)*binaryTree.DataMemoryLocation{
	var fileNameMetaData string = tableName + "_metaData.txt"
	var fileName string  = tableName + ".txt"

	var lineNumber int = SetLineNumber(fileNameMetaData, key)

	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		check(err,"Error opening file for append")
	}

	defer f.Close()
	//	persist data in fileName, need to convert payload to string then convert back when getting

	return &binaryTree.DataMemoryLocation{LineNumber:lineNumber}
}

func AppendFileTableMeta(fileName string, location string, schema globalTypes.TableSchema) {
	finalLocation := location+"/"+fileName
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

	CreateFile(tableFileName,".")
	CreateFile(tableFileNameMetaData,".")

	_, err = f.WriteString("\n")
	check(err, "Faied to write to file")

	// _,err := f.Close()
	// check(err, "Error closing file")
}


