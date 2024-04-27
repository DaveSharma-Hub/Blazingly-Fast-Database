package commandArgs

import (
	"strings"
	"strconv"
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/types"
	"fmt"
)

func argumentValidation(input string)bool{
	var character = []string{"="}
	for i:=0; i<len(character); i++ {
		if !strings.Contains(input,"=") {
			return false
		}
	}
	return true
}

func ParseInput(arguments []string)globalTypes.CommandLineArguments{
	commandLineArgs := globalTypes.CommandLineArguments{CacheMaxSize:10000}
	for i := 0; i<len(arguments) ; i++ {
		isValid := argumentValidation(arguments[i])
		arg := strings.Split(arguments[i],"=")
		if !isValid {
			panic("Invalid arguments")
		}
		var argType string = arg[0]
		var argValue string = arg[1]
		switch argType{
		case "--mc":
			size, err := strconv.ParseInt(argValue, 10, 64)
			if err == nil {
				commandLineArgs.CacheMaxSize = size
			}else{
				panic(err)
			}
		default:
			// do nothing
			fmt.Println("Cache set to default of 10000 items")
		}
	}
	return commandLineArgs
}
