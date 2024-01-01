package commandArgs

import (
	"strings"
	"strconv"
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/types"
)

func ParseInput(arguments []string)globalTypes.CommandLineArguments{
	var cacheSize int64 = 1000
	
	for i := 0; i<len(arguments) ; i++ {
		arg := strings.Split(arguments[i],"=")
		size, err := strconv.ParseInt(arg[1], 10, 64)
		if err == nil {
			cacheSize = size
		}else{
			panic(err)
		}
	}
	return globalTypes.CommandLineArguments{CacheMaxSize: cacheSize}
}
