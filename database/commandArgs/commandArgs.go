package commandArgs

import (
	"strings"
	"strconv"
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/types"
	"fmt"
)

func ParseInput(arguments []string)globalTypes.CommandLineArguments{
	commandLineArgs := globalTypes.CommandLineArguments{CacheMaxSize:10000}
	for i := 0; i<len(arguments) ; i++ {
		arg := strings.Split(arguments[i],"=")
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
